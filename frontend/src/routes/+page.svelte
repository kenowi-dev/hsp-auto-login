<script lang="ts">
    import type { PageProps } from './$types';
    import type {Course} from "$lib/types";

    let { data }: PageProps = $props();


    let selectedSport: string = $state("")
    let courses: Course[] = $state([])
    async function getCourses() {
        const response = await fetch(`/api/courses?sport=${selectedSport}`)
        const json = await response.json()
        console.log(json)
        courses = json as Course[]
    }
</script>

<div class="flex items-center min-h-screen p-6 bg-gray-800">
    <div class="flex-1 max-w-md mx-auto bg-gray-900 rounded-md shadow-md overflow-hidden">
        <div class="py-4 px-6">
            <h2 class="text-2xl font-semibold text-gray-200">Sports Website Login</h2>
            <!-- Rest of the login form -->
            <form id="loginForm">
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

                <!-- Dropdown for courses -->
                <div class="mt-4">
                    <label class="block text-gray-400 text-sm font-bold" for="coursesDropdown">Select Course</label>
                    <select name="courseNumber"
                            id="coursesDropdown"
                            class="mt-1 p-2 w-full border rounded-md bg-gray-700 text-white">
                        <option value="" disabled>Select Course</option>
                        {#each courses as course}
                            <option value={course.id}>{course.number}>{course.number} -- {course.details} -- {course.day} -- {course.time}</option>
                        {/each}

                        <!-- Courses will be populated here dynamically -->
                    </select>
                </div>

                <!-- Dropdown for time slots -->
                <div class="mt-4">
                    <label class="block text-gray-400 text-sm font-bold" for="timeSlotDropdown">Select Time Slot</label>
                    <select name="courseNumber" id="timeSlotDropdown" class="mt-1 p-2 w-full border rounded-md bg-gray-700 text-white">
                        <option value="">Select Time Slot</option>
                        <!-- Time slots will be populated here dynamically -->
                    </select>
                </div>

                <div class="mt-4">
                    <label for="email" class="block text-gray-400 text-sm font-bold">Email</label>
                    <input id="email" type="email" name="email" class="mt-1 p-2 w-full border rounded-md bg-gray-700 text-white">
                </div>

                <div class="mt-4">
                    <label for="password" class="block text-gray-400 text-sm font-bold">Password</label>
                    <input id="password" type="password" name="password" class="mt-1 p-2 w-full border rounded-md bg-gray-700 text-white">
                </div>
                <div class="mt-6 flex flex-row justify-between">
                    <button form="loginForm"
                            class="px-4 py-2 text-white bg-green-500 rounded-md hover:bg-green-600 focus:outline-none focus:shadow-outline-blue active:bg-green-800">
                        See active registration scanners
                    </button>
                    <button form="loginForm"
                            class="px-4 py-2 text-white bg-blue-500 rounded-md hover:bg-blue-600 focus:outline-none focus:shadow-outline-blue active:bg-blue-800">
                        Login
                    </button>
                    <button form="loginForm"
                            class="px-4 py-2 text-white bg-red-500 rounded-md hover:bg-red-600 focus:outline-none focus:shadow-outline-blue active:bg-red-800">
                        dev
                    </button>
                </div>
            </form>
            <div id="registrations" class="mt-4"></div>
        </div>
    </div>
</div>