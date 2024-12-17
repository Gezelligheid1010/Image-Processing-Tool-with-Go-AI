import http from 'k6/http';
import { check, group } from 'k6';

export let options = {
    stages: [
        { duration: '0.5m', target: 10 }, // 前 30 秒增加到 10 个用户
        { duration: '1m', target: 50 }, // 1 分钟保持 50 个用户
        { duration: '0.5m', target: 0 }, // 最后 30 秒逐步减少到 0
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

        check(response, {
            "response should have the correct state": res => res.json()[0].completed === false,
        });
    })
}
