import http from 'k6/http';
import { check, group } from 'k6';
import { Trend } from 'k6/metrics';

const uptimeTrendCheck = new Trend('/GET API uptime');
const todoCreationTrend = new Trend('/POST Create a todo');


export let options = {
    stages: [
        { duration: '0.5m', target: 3 }, // simulate ramp-up of traffic from 0 to 3 users over 0.5 minutes.
        // { duration: '0.5m', target: 4 }, 
        // { duration: '0.5m', target: 0 }, // ramp-down to 0 users
    ],
};

export default function () {
    group('API uptime check', () => {
        const response = http.get('https://todo-app-barkend.herokuapp.com/todos/');
        uptimeTrendCheck.add(response.timings.duration);
        check(response, {
            "status code should be 200": res => res.status === 200,
        });
    });

    let todoID;
    group('Create a Todo', () => {
        const response = http.post('https://todo-app-barkend.herokuapp.com/todos/',
            { "task": "write k6 tests" }
        );
        todoCreationTrend.add(response.timings.duration);
        todoID = response.json()._id;
        check(response, {
            "status code should be 200": res => res.status === 200,
        });
        check(response, {
            "response should have created todo": res => res.json().completed === false,
        });
    })
}
