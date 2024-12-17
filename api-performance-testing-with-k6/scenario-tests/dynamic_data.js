import { b64encode } from 'k6/encoding';

const base64Image = ""

export function generateUniqueUsername() {
    const randomPart = Math.random().toString(36).substring(2, 10); // 随机部分
    const timestamp = Date.now().toString(36); // 时间戳部分
    // console.log('name:',`user_${randomPart}${timestamp}`)
    return `user_${randomPart}${timestamp}`;
}

export function generateUniqueEmail() {
    return `test_${Math.random().toString(36).substring(2, 15)}@example.com`;
}

// 模拟动态生成 Base64 图片
export function generateImage(base64Image) {
    const randomData = Array.from({ length: 10 }, () =>
        Math.floor(Math.random() * 256)
    );
    const randomBase64 = b64encode(String.fromCharCode(...randomData));
    // console.log("randomBase64；",randomBase64.length)
    // console.log("base64Image；",base64Image.length)
    return base64Image.slice(0, -randomBase64.length) + randomBase64; // 替换部分数据
}


function randomChoice(arr) {
    return arr[Math.floor(Math.random() * arr.length)];
}

export function generateRandomPrompt() {
    const subjects = ["cat", "robot", "wizard", "forest", "cityscape", "mountain"];
    const adjectives = ["mystical", "vibrant", "dark", "dreamy", "futuristic"];
    const backgrounds = ["a dense jungle", "a futuristic city", "an ancient temple"];
    const times = ["sunset", "dawn", "a rainy afternoon", "a snowy night"];
    const styles = ["watercolor", "oil painting", "pixel art", "cyberpunk"];
    const details = ["soft lighting", "vivid colors", "intricate patterns", "minimalist design"];

    return `A ${randomChoice(adjectives)} ${randomChoice(subjects)} in ${randomChoice(backgrounds)} during ${randomChoice(times)}, drawn in ${randomChoice(styles)} style with ${randomChoice(details)}.`;
}