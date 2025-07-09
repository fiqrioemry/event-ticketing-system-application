// src/lib/types/event.ts
export interface Event {
	id: string;
	name: string;
	price: number;
	description: string;
	image: string;
}

export interface CreateEventDto {
	name: string;
	price: number;
	description: string;
	image: string;
}

export interface UpdateEventDto extends Partial<CreateEventDto> {
	id: string;
}
