// utils/export.go
package utils

import (
	"encoding/csv"
	"fmt"
	"net/http"
	"reflect"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
)

func ExportCSV(c *gin.Context, filename string, records any) {
	v := reflect.ValueOf(records)
	if v.Kind() != reflect.Slice {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "invalid data format for CSV export"})
		return
	}

	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "text/csv; charset=utf-8")
	c.Header("Cache-Control", "no-cache")

	writer := csv.NewWriter(c.Writer)

	// headers
	if v.Len() > 0 {
		first := v.Index(0)
		var headers []string
		for i := 0; i < first.NumField(); i++ {
			headers = append(headers, first.Type().Field(i).Tag.Get("json"))
		}
		writer.Write(headers)
	}

	// rows
	for i := 0; i < v.Len(); i++ {
		val := v.Index(i)
		var row []string
		for j := 0; j < val.NumField(); j++ {
			row = append(row, toString(val.Field(j).Interface()))
		}
		writer.Write(row)
	}

	writer.Flush()
}

func toString(v interface{}) string {
	return fmt.Sprintf("%v", v)
}

func ExportPDF(c *gin.Context, filename string, records interface{}) {
	v := reflect.ValueOf(records)
	if v.Kind() != reflect.Slice || v.Len() == 0 {
		c.JSON(500, gin.H{"error": "Invalid or empty data for PDF export"})
		return
	}

	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.SetFont("Arial", "", 10)
	pdf.AddPage()

	first := v.Index(0)
	numCols := first.NumField()
	var headers []string
	var colWidths []float64

	for i := 0; i < numCols; i++ {
		field := first.Type().Field(i)
		fieldName := field.Tag.Get("json")
		headers = append(headers, fieldName)

		switch fieldName {
		case "orderId", "paymentId", "userId", "withdrawalId":
			colWidths = append(colWidths, 60)
		case "email":
			colWidths = append(colWidths, 55)
		case "amount", "status":
			colWidths = append(colWidths, 30)
		case "createdAt", "approvedAt", "refundedAt":
			colWidths = append(colWidths, 45)
		default:
			colWidths = append(colWidths, 40)
		}
	}

	for i := range numCols {
		pdf.CellFormat(colWidths[i], 8, headers[i], "1", 0, "", false, 0, "")
	}
	pdf.Ln(-1)

	for i := range v.Len() {
		row := v.Index(i)
		for j := 0; j < numCols; j++ {
			cellValue := fmt.Sprintf("%v", row.Field(j).Interface())
			pdf.CellFormat(colWidths[j], 8, truncate(cellValue, int(colWidths[j]/2.5)), "1", 0, "", false, 0, "")
		}
		pdf.Ln(-1)
	}

	c.Header("Content-Disposition", "attachment; filename="+filename)
	c.Header("Content-Type", "application/pdf")
	c.Header("Cache-Control", "no-cache")

	if err := pdf.Output(c.Writer); err != nil {
		c.JSON(500, gin.H{"error": "failed to generate PDF"})
	}
}

func truncate(s string, max int) string {
	if len(s) > max {
		return s[:max-3] + "..."
	}
	return s
}
