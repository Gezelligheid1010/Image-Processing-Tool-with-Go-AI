import http from 'k6/http';
import { check, group } from 'k6';
import { stagingConfig } from '../environmentConfig.js';

export const options = {
    scenarios: {
      example_scenario: {
        env: stagingConfig(),
        executor: 'shared-iterations',
      }
    }
  };

export default function () {
    group('API uptime check - Staging', () => {
        const response = http.get(`${__ENV.BASE_URL}`);
        check(response, {
            "status code should be 200": res => res.status === 200,
        });
    });

    let todoID;
    group('Create a Todo', () => {
        const response = http.post(`${__ENV.BASE_URL}`, 
        {"task": "write k6 tests"}
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
        const response = http.get(`${__ENV.BASE_URL}${todoID}`
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

    group('delete all Todos', () => {
        let response = http.del(`${__ENV.BASE_URL}`
        );

        check(response, {
            "status code should be 200": res => res.status === 200,
        });
    })
}
