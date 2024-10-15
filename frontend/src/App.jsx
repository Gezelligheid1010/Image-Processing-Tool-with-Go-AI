import React, { useState,useRef } from 'react';
import Navbar from './components/Navbar';
import ImageUpload from './components/ImageUpload';
import Footer from './components/Footer';
import DrawingCanvas from './components/DrawingCanvas';

function App() {
    const [uploadedImage, setUploadedImage] = useState(null);
    const [processedImage, setProcessedImage] = useState(null);
    const [prompt, setPrompt] = useState(''); // 用于存储用户输入的 prompt
    const drawingCanvasRef = useRef();


    const handleProcessImage = async () => {
        const canvasDataURL = await drawingCanvasRef.current.getCanvasDataURL();
        // console.log('Canvas URL:', canvasDataURL);

        const response = await fetch('http://localhost:8080/api/process', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({
                prompt: prompt,
                // init_image: "https://raw.githubusercontent.com/CompVis/stable-diffusion/main/data/inpainting_examples/overture-creations-5sI6fQgYIuo.png",
                init_image: canvasDataURL,
                width: '512',
                height: '512',
                num_inference_steps: '30',
                guidance_scale: 7.5,
            }),
        });

        // const result = await response.json();
        const responseText = await response.text();
        console.log('Raw response:', responseText);

        // 确保返回的数据是 JSON 格式
        const result = JSON.parse(responseText);
        setProcessedImage(result.processedImage);
    };


    const handleImageUpload = (file) => {
        const imageUrl = URL.createObjectURL(file);
        setUploadedImage(imageUrl);
    };

    const handleSaveImage = () => {
        if (processedImage) {
            const link = document.createElement('a');
            link.href = processedImage;
            link.download = 'processed-image.png';
            link.click();
        }
    };

    return (
        <div className="App">
            <Navbar/>


            <div className="app-container">
                <div className="left-panel">
                    <h2>手绘图像</h2>
                    <DrawingCanvas ref={drawingCanvasRef} uploadedImage={uploadedImage}/>
                    <ImageUpload onImageUpload={handleImageUpload}/>
                </div>

                <div className="separator"></div>

                <div className="right-panel">
                    <h2>AI 处理图像</h2>

                    <div className="processed-image-area">
                        {processedImage && <img src={processedImage} alt="Processed"/>}
                    </div>

                    {/* 用户输入 prompt 的输入框 */}
                    <div className="prompt-input">
                        <input
                            type="text"
                            placeholder="请输入描述..."
                            value={prompt}
                            onChange={(e) => setPrompt(e.target.value)}
                        />
                    </div>
                    <button className="ai-process-btn" onClick={handleProcessImage}>
                        {/*onClick={() => document.querySelector('.ai-process-btn').click()}>*/}
                        AI Process Image
                    </button>
                    <button className="save-btn" onClick={handleSaveImage}>Save Image</button>
                    {/* Save button */}
                </div>
            </div>
            <Footer/>
        </div>
    );
}

export default App;
