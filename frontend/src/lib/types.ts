export type Sport = {
    name: string;
    href: string;
    inFlexiCard: boolean;
    extraInfo: string;
}

export type Course = {
    number: string;
    details: string;
    day: string;
    time: string;
    id: string;
    location: string;
    management: string;
    price: string;
    state: CourseState;
    bookingID: string;
}

export type CourseState = "Vormerkliste" | "Warteliste"


export type CourseDate = {
    start: Date
    end: Date
    duration: string
}