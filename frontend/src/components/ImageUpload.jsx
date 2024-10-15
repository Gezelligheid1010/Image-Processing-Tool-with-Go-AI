import React, { useState } from 'react';

function ImageUpload({ onImageUpload }) {
    const [selectedImage, setSelectedImage] = useState(null);

    const handleImageChange = (e) => {
        const file = e.target.files[0];
        // setSelectedImage(URL.createObjectURL(file));
        onImageUpload(file);
    };

    return (
        <div className="my-8 p-4 border-dashed border-2 border-gray-300 rounded-lg text-center">
            <input type="file" accept="image/*" onChange={handleImageChange} className="hidden" id="upload-input"/>
            <label htmlFor="upload-input" className="cursor-pointer text-blue-600">Click to Upload Image</label>
            {/*{selectedImage && <img src={selectedImage} alt="Preview" className="mx-auto mt-4 w-1/2"/>}*/}
        </div>
    );
}

export default ImageUpload;
