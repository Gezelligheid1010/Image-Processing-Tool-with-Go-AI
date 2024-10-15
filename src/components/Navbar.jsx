import React from 'react';

function Navbar() {
    return (
        <nav className="bg-blue-600 p-4">
            <div className="container mx-auto flex justify-between items-center">
                <h1 className="text-white text-2xl font-bold">AI Hand-Drawn Image</h1>
                <div>
                    <a href="#home" className="text-white mx-4">Home</a>
                    <a href="#upload" className="text-white mx-4">AI Processing</a>
                    <a href="#gallery" className="text-white mx-4">My Works</a>
                </div>
            </div>
        </nav>
    );
}

export default Navbar;
