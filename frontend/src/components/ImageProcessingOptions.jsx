import React from 'react';

function ImageProcessingOptions() {
    return (
        <div className="my-4 p-4">
            <h3 className="text-xl font-bold mb-4">Choose AI Style</h3>
            <select className="border border-gray-300 p-2 rounded-lg">
                <option value="sketch">Sketch</option>
                <option value="watercolor">Watercolor</option>
                <option value="oilpainting">Oil Painting</option>
            </select>
        </div>
    );
}

export default ImageProcessingOptions;
