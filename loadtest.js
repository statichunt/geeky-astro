import http from 'k6/http';
import {check, sleep} from 'k6';

export const options = {
    stages: [
        {duration: '10s', target: 100},
        {duration: '5m', target: 1300},
        {duration: '20s', target: 0},
    ],
}

export default function () {
    const pages = ['/', '/categories/rugby/', '/posts/springboks-vs-scotland/', '/posts/wales-triumph-over-fiji-in-thrilling-rugby-world-cup-opener/']

    for (const page of pages) {
        const res = http.get(`http://localhost${page}`);
        check(res, {
            'status was 200': (r) => r.status == 200,
            'transaction time OK': (r) => r.timings.duration < 200,
        });
        sleep(1);
    }
}