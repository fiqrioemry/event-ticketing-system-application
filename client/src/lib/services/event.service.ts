// lib/services/event.service.ts
import qs from 'qs';
import { buildFormData } from '$lib/utils/formatter';
import { publicInstance, authInstance } from '$lib/services/client';
import type { EventQueryParams, CreateEventRequest, UpdateEventRequest } from '$lib/types/api';

// GET /api/events?search=&status=&page=&limit=
export const getAllEvents = async (params?: EventQueryParams) => {
	const queryString = qs.stringify(params, { skipNulls: true });
	const res = await publicInstance.get(`/events?${queryString}`);
	return res.data;
};

// GET /api/events/:id
export const getEventById = async (id: string) => {
	const res = await publicInstance.get(`/events/${id}`);
	return res.data;
};

// GET /api/events/:id/tickets
export const getTicketsByEventId = async (id: string) => {
	const res = await publicInstance.get(`/events/${id}/tickets`);
	return res.data;
};

// POST /api/events (Admin only)
export const createEvent = async (data: CreateEventRequest) => {
	const formData = buildFormData(data);
	const res = await authInstance.post('/events', formData);
	return res.data;
};

// PUT /api/events/:id (Admin only)
export const updateEvent = async (id: string, data: UpdateEventRequest) => {
	const formData = buildFormData(data);
	const res = await authInstance.put(`/events/${id}`, formData);
	return res.data;
};

// DELETE /api/events/:id (Admin only)
export const deleteEvent = async (id: string) => {
	const res = await authInstance.delete(`/events/${id}`);
	return res.data;
};
