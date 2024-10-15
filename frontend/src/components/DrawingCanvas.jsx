import React, { useRef, useState, useEffect, forwardRef, useImperativeHandle } from 'react';

const DrawingCanvas = forwardRef(({ uploadedImage }, ref) => {
    const canvasRef = useRef(null);
    const [isDrawing, setIsDrawing] = useState(false);
    const [ctx, setCtx] = useState(null);

    // 设置
    const [color, setColor] = useState('#000000'); // 默认黑色
    const [brushSize, setBrushSize] = useState(5); // 默认画笔大小
    const [tool, setTool] = useState('brush'); // 默认工具为画笔

    const [scale, setScale] = useState(1); // 图像缩放比例
    const [rotation, setRotation] = useState(0); // 图像旋转角度

    useEffect(() => {
        const canvas = canvasRef.current;
        const context = canvas.getContext('2d');
        setCtx(context);

        // 如果上传了图片，将图片显示在画布上
        if (uploadedImage) {
            drawImageOnCanvas(uploadedImage, context);
        }
    }, [uploadedImage, scale, rotation]);

    // 使用 useImperativeHandle 来暴露方法给父组件
    useImperativeHandle(ref, () => ({
        getCanvasDataURL: () => {
            return canvasRef.current.toDataURL('image/png'); // 返回Base64格式的图像数据
        }
    }));

    const drawImageOnCanvas = (imageSrc, context) => {
        const image = new Image();
        image.src = imageSrc;
        image.onload = () => {
            context.clearRect(0, 0, canvasRef.current.width, canvasRef.current.height);

            // 根据画布和图像的比例动态调整大小
            const aspectRatio = image.width / image.height;
            let drawWidth = canvasRef.current.width;
            let drawHeight = drawWidth / aspectRatio;

            if (drawHeight > canvasRef.current.height) {
                drawHeight = canvasRef.current.height;
                drawWidth = drawHeight * aspectRatio;
            }

            // 设置旋转和缩放
            context.save();
            context.translate(canvasRef.current.width / 2, canvasRef.current.height / 2);
            context.rotate((rotation * Math.PI) / 180);
            context.scale(scale, scale);
            context.drawImage(image, -drawWidth / 2, -drawHeight / 2, drawWidth, drawHeight);
            context.restore();
        };
    };

    const handleZoomIn = () => setScale(scale + 0.1);
    const handleZoomOut = () => setScale(scale - 0.1);
    const handleRotateLeft = () => setRotation(rotation - 15);
    const handleRotateRight = () => setRotation(rotation + 15);
    const handleReset = () => {
        setScale(1);
        setRotation(0);
    };

    const clearCanvas = () => {
        const ctx = canvasRef.current.getContext('2d');
        ctx.clearRect(0, 0, canvasRef.current.width, canvasRef.current.height);
    };

    const handleToolChange = (selectedTool) => {
        setTool(selectedTool);
    };

    const startDrawing = (e) => {
        setIsDrawing(true);
        ctx.beginPath();
        ctx.moveTo(e.nativeEvent.offsetX, e.nativeEvent.offsetY);
    };

    const draw = (e) => {
        if (!isDrawing) return;
        const ctx = canvasRef.current.getContext('2d');
        ctx.lineTo(e.nativeEvent.offsetX, e.nativeEvent.offsetY);
        ctx.lineWidth = brushSize;
        ctx.lineCap = 'round';

        if (tool === 'brush') {
            ctx.strokeStyle = color;
        } else if (tool === 'eraser') {
            ctx.strokeStyle = '#ffffff'; // 橡皮擦使用白色
        }
        ctx.stroke();
    };

    const stopDrawing = () => {
        setIsDrawing(false);
        ctx.closePath();
    };

    return (
        <div style={{ textAlign: 'center' }}>
            <div style={{ marginBottom: '10px' }}>
                <label>
                    选择颜色:
                    <input
                        type="color"
                        value={color}
                        onChange={(e) => setColor(e.target.value)}
                    />
                </label>
                <label style={{ marginLeft: '10px' }}>
                    笔刷大小:
                    <input
                        type="range"
                        min="1"
                        max="20"
                        value={brushSize}
                        onChange={(e) => setBrushSize(e.target.value)}
                    />
                </label>
                <button onClick={() => handleToolChange('brush')} style={{ marginLeft: '10px' }}>
                    画笔
                </button>
                <button onClick={() => handleToolChange('eraser')} style={{ marginLeft: '5px' }}>
                    橡皮擦
                </button>
                <button onClick={clearCanvas} style={{ marginLeft: '5px' }}>
                    清除画布
                </button>
            </div>
            <div className="controls">
                <button onClick={handleZoomIn}>Zoom In</button>
                <button onClick={handleZoomOut}>Zoom Out</button>
                <button onClick={handleRotateLeft}>Rotate Left</button>
                <button onClick={handleRotateRight}>Rotate Right</button>
                <button onClick={handleReset}>Reset</button>
            </div>
            <canvas
                ref={canvasRef}
                width={600}
                height={400}
                onMouseDown={startDrawing}
                onMouseMove={draw}
                onMouseUp={stopDrawing}
                onMouseLeave={stopDrawing}
                style={{ border: '1px solid black' }}
            />
        </div>
    );
});

export default DrawingCanvas;
