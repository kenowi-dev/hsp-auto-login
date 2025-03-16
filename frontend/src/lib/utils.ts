export function formatNanoseconds(nanoseconds: number): string {
    const seconds = Math.floor(nanoseconds / 1e9); // Convert nanoseconds to seconds
    const hours = Math.floor(seconds / 3600); // Calculate hours
    const minutes = Math.floor((seconds % 3600) / 60); // Calculate minutes
    const remainingSeconds = seconds % 60; // Calculate remaining seconds

    // Format as HH:MM:SS, ensuring two digits for each part
    return `${String(hours).padStart(2, '0')}:${String(minutes).padStart(2, '0')}:${String(remainingSeconds).padStart(2, '0')}`;
}


export function formatDateTime(date: Date) {
    let f = new Intl.DateTimeFormat('de', {
        year: 'numeric',
        month: 'short',
        day: 'numeric',
        weekday: 'short',
        hour: '2-digit',
        hour12: false,
        minute: '2-digit',
    });
    return f.format(date);
}

export function formatDate(date: Date) {
    let f = new Intl.DateTimeFormat('de', {
        year: 'numeric',
        month: 'short',
        day: 'numeric',
        weekday: 'short'
    });
    return f.format(date);
}

export function formatTime(date: Date) {
    let f = new Intl.DateTimeFormat('de', {
        hour: '2-digit',
        hour12: false,
        minute: '2-digit'
    });
    return f.format(date);
}

export function parseTimeRange(timeRange: string): [ start: Date, end: Date ] {
    const [startTime, endTime] = timeRange.split('-'); // Split the time range into start and end
    const today = new Date(); // Get today's date

    // Set the hours and minutes for the start time
    const start = new Date(today);
    const [startHour, startMinute] = startTime.split(':').map(Number);
    start.setHours(startHour, startMinute, 0, 0); // Set the start time with hours, minutes, seconds

    // Set the hours and minutes for the end time
    const end = new Date(today);
    const [endHour, endMinute] = endTime.split(':').map(Number);
    end.setHours(endHour, endMinute, 0, 0); // Set the end time with hours, minutes, seconds

    return [ start, end ];
}

export function parseDateWithTime(dateString: string, timeRange: string): [ start: Date, end: Date ] {
    const startDate = new Date(dateString);
    const endDate = new Date(dateString);

    const [startTime, endTime] = parseTimeRange(timeRange)

    // Set the hours and minutes for the parsed date
    startDate.setTime(startTime.getTime())

    endDate.setTime(endTime.getTime())

    return [ startDate, endDate ];
}