const dateOptions: Intl.DateTimeFormatOptions = {
	day: "2-digit",
	month: "2-digit",
	year: "numeric",
}

export function formatDate(date?: Date): string {
	
	if (date) {
		return date.toLocaleDateString("ru", dateOptions)
	}

	return "";
}