<script lang="ts">
    import type {PageProps} from './$types';
    import type {Course, CourseDate} from "$lib/types";
    import {formatDateTime, formatNanoseconds, parseDateWithTime} from "$lib/utils";
    import {formatTime} from "$lib/utils.js";

    let {data}: PageProps = $props();


    let error: string | undefined = $state()
    let selectedSport = $state("")
    let selectedCourseIdx = $state(0)
    let selectedDateIdx = $state(0)

    let courses: Course[] = $state([])
    let dates: CourseDate[] = $state([])
    let email: string | undefined = $state()
    let password: string | undefined = $state()


    async function getCourses() {
        const response = await fetch(`/api/courses?sport=${selectedSport}`);
        const json = await response.json();
        courses = json as Course[];
        selectedCourseIdx = courses.length - 1;
        await getCourseDates();
    }

    async function getCourseDates() {
        const course = courses[selectedCourseIdx]
        const response = await fetch(`/api/coursesDates`, {
            method: "POST",
            body: JSON.stringify(course)
        })
        const json = await response.json() as { date: string; duration: number }[]
        console.log(json)
        dates = json.map(value => {
            const when = parseDateWithTime(value.date, course.time)
            return {
                start: when[0],
                end: when[1],
                duration: formatNanoseconds(value.duration)
            }
        })
        console.log($state.snapshot(dates))
    }

    async function register() {

        const response = await fetch(`/api/register`, {
            method: "POST",
            body: JSON.stringify({
                course: courses[selectedCourseIdx],
                sport: selectedSport,
                date: dates[selectedDateIdx].start,
                email: email,
                password: password
            })
        })
        if (!response.ok) {
            error = await response.text()
        } else {
            window.alert("YEY")
        }
    }
</script>

<div class="flex items-center min-h-screen p-6 bg-gray-800">
    <div class="flex-1 max-w-md mx-auto bg-gray-900 rounded-md shadow-md overflow-hidden">
        <div class="py-4 px-6">
            <h2 class="text-2xl font-semibold text-gray-200">Sports Website Login</h2>
            <div>
                <!-- Dropdown for available sports -->
                <div class="mt-4">
                    <label class="block text-gray-400 text-sm font-bold" for="sportDropdown">Select Sport</label>
                    <select bind:value={selectedSport}
                            onchange={getCourses}
                            class="mt-1 p-2 w-full border rounded-md bg-gray-700 text-white">
                        <option value="" disabled>Select Sport</option>
                        {#await data.sports then sports}
                            {#each sports as sport}
                                <option value={sport.name}>{ sport.name }</option>
                            {/each}
                        {/await}
                    </select>
                </div>

                {#if courses.length > 0}
                    <!-- Dropdown for courses -->
                    <div class="mt-4">
                        <label class="block text-gray-400 text-sm font-bold" for="coursesDropdown">Select Course</label>
                        <select bind:value={selectedCourseIdx}
                                onchange={getCourseDates}
                                class="mt-1 p-2 w-full border rounded-md bg-gray-700 text-white">
                            <option value="" disabled>Select Course</option>
                            {#each courses as course, i}
                                <option value={i}>
                                    {course.number} -- {course.details} -- {course.day}-- {course.time}
                                </option>
                            {/each}
                        </select>
                    </div>
                {/if}

                {#if dates.length > 0}
                    <!-- Dropdown for time slots -->
                    <div class="mt-4">
                        <label class="block text-gray-400 text-sm font-bold" for="timeSlotDropdown">Select Time
                            Slot</label>
                        <select bind:value={selectedDateIdx}
                                class="mt-1 p-2 w-full border rounded-md bg-gray-700 text-white">
                            <option value="" disabled>Select Time Slot</option>
                            {#each dates as date, i}
                                <option value={i}>{formatDateTime(date.start)} - {formatTime(date.end)}</option>
                            {/each}
                        </select>
                    </div>


                    <div class="mt-4">
                        <label for="email" class="block text-gray-400 text-sm font-bold">Email</label>
                        <input id="email" type="email"
                               bind:value={email}
                               class="mt-1 p-2 w-full border rounded-md bg-gray-700 text-white">
                    </div>

                    <div class="mt-4">
                        <label for="password" class="block text-gray-400 text-sm font-bold">Password</label>
                        <input id="password" type="password"
                               bind:value={password}
                               class="mt-1 p-2 w-full border rounded-md bg-gray-700 text-white">
                    </div>
                {/if}

                {#if error}
                    <span class="text-red-500 font-bold">{error}</span>
                {/if}

                {#if dates.length > 0}
                    <div class="mt-6 flex flex-row justify-between">
                        <button
                                class="px-4 py-2 text-white bg-green-500 rounded-md hover:bg-green-600 focus:outline-none focus:shadow-outline-blue active:bg-green-800">
                            See active registration scanners
                        </button>
                        <button onclick={register}
                                class="px-4 py-2 text-white bg-blue-500 rounded-md hover:bg-blue-600 focus:outline-none focus:shadow-outline-blue active:bg-blue-800">
                            Register
                        </button>
                    </div>
                {/if}
            </div>
            <div id="registrations" class="mt-4"></div>
        </div>
    </div>
</div>