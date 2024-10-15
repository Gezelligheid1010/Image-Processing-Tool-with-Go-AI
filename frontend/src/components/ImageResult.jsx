import React from 'react';

function ImageResult({ originalImage, processedImage }) {
    return (
        <div className="my-8 flex justify-between items-center">
            <div>
                <h4 className="text-lg font-bold">Original Image</h4>
                {originalImage && <img src={originalImage} alt="Original" className="w-64 h-64 border" />}
            </div>
            <div>
                <h4 className="text-lg font-bold">AI Processed Image</h4>
                {processedImage ? <img src={processedImage} alt="Processed" className="w-64 h-64 border" /> : <p>Processing...</p>}
            </div>
        </div>
    );
}

export default ImageResult;
