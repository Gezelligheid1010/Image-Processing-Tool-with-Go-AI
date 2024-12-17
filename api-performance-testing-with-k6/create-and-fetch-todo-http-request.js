import http from 'k6/http';
import { check, group } from 'k6';

export let options = {
    stages: [
        { duration: '0.5m', target: 3 }, // simulate ramp-up of traffic from 1 to 100 users over 5 minutes.
        // { duration: '0.5m', target: 4 }, // stay at 100 users for 10 minutes
        // { duration: '0.5m', target: 0 }, // ramp-down to 0 users
    ],
};

export default function () {
    group('API uptime check', () => {
        const response = http.get('https://todo-app-barkend.herokuapp.com/todos/');
        check(response, {
            "status code should be 200": res => res.status === 200,
        });
    });

    let todoID;
    group('Create a Todo', () => {
        const response = http.post('https://todo-app-barkend.herokuapp.com/todos/',
            { "task": "write k6 tests" }
        );
        todoID = response.json()._id;
        check(response, {
            "status code should be 200": res => res.status === 200,
        });
        check(response, {
            "response should have created todo": res => res.json().completed === false,
        });
    })

    group('get a todo item', () => {
        const response = http.get(`https://todo-app-barkend.herokuapp.com/todos/${todoID}`
        );
        check(response, {
            "status code should be 200": res => res.status === 200,
        });
        check(response, {
            "response should have the created todo": res => res.json()[0]._id === todoID,
        });
        console.log(JSON.stringify(response.json()[0]));

        check(response, {
            "response should have the correct state": res => res.json()[0].completed === false,
        });
    })
}
