import http from 'k6/http';
import { sleep, check } from 'k6';
import {generateImage, generateRandomPrompt, generateUniqueUsername,generateUniqueEmail} from "./dynamic_data.js";

// 配置接口的基本 URL
const BASE_URL = 'http://localhost:8080/api/v1';

// 读取 Base64 数据
const base64Image = open("image_base64.txt");

// 定义一个变量来跟踪是否正在刷新 token
let refreshing = false;

export const options = {
    scenarios: {
        my_test: {
            executor: 'constant-arrival-rate',
            rate: 5, // 每秒发送 10 个请求
            duration: '5m', // 持续 5 分钟
            preAllocatedVUs: 5, // 提前分配的虚拟用户数
            maxVUs: 10, // 最大虚拟用户数
        },
    },
};

// k6 入口函数
export default function () {
    // 用户注册信息
    const USER_CREDENTIALS = {
        avatar: `data:image/png;base64,${generateImage(base64Image)}`,
        username: generateUniqueUsername(),
        password: 'password123',
        confirm_password: 'password123',
        email: generateUniqueEmail(),
    };
    // 1. 用户注册
    const res = http.post(`${BASE_URL}/signup`, JSON.stringify(USER_CREDENTIALS), {
        headers: { 'Content-Type': 'application/json' },
    });
    check(res, { 'signup succeeded': (r) => r.status === 200 });

    // console.log('signup Response:', JSON.stringify(res.json(), null, 2));

    // 2. 用户登录
    const loginRes = http.post(`${BASE_URL}/login`, JSON.stringify(USER_CREDENTIALS), {
        headers: { 'Content-Type': 'application/json' },
    });

    check(loginRes, {
        'Login succeeded': (r) => r.status === 200,
    });

    // Debug: 打印完整响应内容
    // console.log('Login Response:', JSON.stringify(loginRes.json(), null, 2));
    // console.log('Login Response:', loginRes);

    const accessToken = loginRes.json('data.access_token');
    const refreshTokenValue = loginRes.json('data.refresh_token');
    if (!accessToken || !refreshTokenValue) {
        throw new Error('Login failed, tokens not received');
    }
    // sleep(1);

    // 3. 创建分组
    const genCategory = {
        cover: `data:image/png;base64,${generateImage(base64Image)}`,
        category_name: `category_${Math.random().toString(36).substring(2, 15)}`,
        description: generateRandomPrompt(),
    };
    const createGroupRes = apiRequest('POST', `${BASE_URL}/createCategory`, genCategory, accessToken);
    check(createGroupRes, { 'Create group succeeded': (r) => r.status === 200 });
    // sleep(3);

    // 4. 浏览分组
    const groupRes = apiRequest('GET', `${BASE_URL}/category`, null, accessToken);
    check(groupRes, { 'Browse groups succeeded': (r) => r.status === 200 });

    const groups = groupRes.json('data');

    const groupId = groups[0].category_id;

    // 5. 浏览作品
    const worksRes = apiRequest('GET', `${BASE_URL}/categoryDetail/${groupId}`, null, accessToken);
    check(worksRes, { 'Browse works succeeded': (r) => r.status === 200 });

    // 6. 提交图像处理任务
    const imageTaskPayload = {
        ori_image: `data:image/png;base64,${generateImage(base64Image)}`,
        category_id: groupId,
        prompt: generateRandomPrompt(),
    };

    const processRes = apiRequest('POST', `${BASE_URL}/processImage`, imageTaskPayload, accessToken);
    check(processRes, { 'Process task submitted': (r) => r.status === 200 });

    // 8. 上传作品
    const newWork = {
        work_image: `data:image/png;base64,${generateImage(base64Image)}`,
        // work_image: `data:image/png;base64,${genImage}`,
        category_id: groupId,
        prompt: generateRandomPrompt(),
    };

    const uploadRes = apiRequest('POST', `${BASE_URL}/uploadWork`, newWork, accessToken);
    check(uploadRes, { 'Upload work succeeded': (r) => r.status === 200 });

    // console.log(`Uploaded work with response: ${uploadRes.body}`);

    // 模拟用户操作的间隔
    sleep(1);
}

// 刷新 Token 的逻辑
function refreshToken(refreshToken) {
    const refreshRes = http.post(`${BASE_URL}/refresh_token`, JSON.stringify({ refresh_token: refreshToken }), {
        headers: { 'Content-Type': 'application/json' },
    });

    check(refreshRes, {
        'Token refreshed': (r) => r.status === 200,
    });

    const newToken = refreshRes.json('access_token');
    if (!newToken) {
        throw new Error('Failed to refresh token');
    }

    return newToken;
}

// API 请求方法
function apiRequest(method, url, body, token) {
    const headers = {
        'Content-Type': 'application/json',
        Authorization: `Bearer ${token}`,
    };

    let res;
    if (method === 'GET') {
        res = http.get(url, { headers });
    } else if (method === 'POST') {
        res = http.post(url, JSON.stringify(body), { headers });
    }

    // 如果遇到 401 错误，且没有正在刷新 token
    if (res.status === 401 && !refreshing) {
        console.log('Access token expired, refreshing...');

        // 设置 refreshing 为 true，避免并发请求时重复刷新
        refreshing = true;

        // 刷新 token
        const newToken = refreshToken(refreshTokenValue);

        // 刷新完成后，重新发送请求
        res = apiRequest(method, url, body, newToken);

        // 刷新完成，重置 refreshing 状态
        refreshing = false;
    }

    return res;
}

