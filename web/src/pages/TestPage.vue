<template>
    <app-layout>
        <div class="py-5 text-center">
            <h2>Static Test Page</h2>
            <p class="lead">
                This is a static test page
            </p>
        </div>
        <br><br>
    </app-layout>
</template>

<script>
    import AppLayout from '../layouts/App.vue'

    import {getCalendar} from '../repositories/trainings'
    import {getUserRole, Trainer} from "../repositories/user";
    import {formatDateTime} from "../date";

    export default {
        components: {
            AppLayout,
        },
        data: function () {
            return {
                'calendar': null,
                'isTrainer': null,
                'userRole': null,
            }
        },
        mounted() {
            let self = this
            getCalendar(function (calendar) {
                self.calendar = calendar
            })
            this.isTrainer = getUserRole() === Trainer;
            this.userRole = getUserRole()
        },
        methods: {
            formatDateTime,
        },
    }
</script>

<style scoped>
    h3 {
        margin: 40px 0 0;
    }

    ul {
        list-style-type: none;
        padding: 0;
    }

    li {
        display: inline-block;
        margin: 0 10px;
    }

    a {
        color: #42b983;
    }

    body {
        background-color: #f5f5f5;
    }

    .old-date {
        text-decoration: line-through;
        opacity: 0.5;
    }
</style>
