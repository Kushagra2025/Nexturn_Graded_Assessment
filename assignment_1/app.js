const fs = require('fs');

function readProducts() {
    try {
        const data = fs.readFileSync('data.json', 'utf-8');
        jsonData = JSON.parse(data);
        console.log(jsonData);
    } catch (error) {
        console.error(error);
    }
}

function addProduct(Product) {
    try {
        const data = fs.readFileSync('data.json', 'utf-8');
        const products = JSON.parse(data);
        if (Product && Product.id && Product.name && Product.category && Product.price !== undefined) {
            if (typeof Product.available === 'boolean') {
                products.push(Product);
                const jsonData = JSON.stringify(products);
                fs.writeFileSync('data.json', jsonData);
                console.log("Product added successfully.");
            }
        } else {
            console.error("Invalid product data. Ensure all fields are provided.");
        }
    } catch (error) {
        console.error(error);
    }
}

function updateProduct(productId, updatedFields) {
    try {
        const data = fs.readFileSync('data.json', 'utf-8');
        const products = JSON.parse(data);

        const productIndex = products.findIndex(p => p.id === productId);
        
        if (productIndex !== -1) {
            products[productIndex] = { ...products[productIndex], ...updatedFields };
            const jsonData = JSON.stringify(products);
            fs.writeFileSync('data.json', jsonData);
            console.log("Product updated successfully.");
        } else {
            console.error("Product not found.");
        }
    } catch (error) {
        console.error(error);
    }
}

function filterAvailable() {
    try {
        const data = fs.readFileSync('data.json', 'utf-8');
        const products = JSON.parse(data);
        availableProducts = products.filter(product => product.available);
        console.log(availableProducts);
    } catch (error) {
        console.error(error);
    }
}

function filterbyCategory(category) {
    try {
        const data = fs.readFileSync('data.json', 'utf-8');
        const products = JSON.parse(data);
        productByCategory = products.filter(product => product.category === category);
        console.log(productByCategory);
    } catch (error) {
        console.error(error);
    }
}


readProducts();

Product = {
    id: 4,
    name: "Badminton Racket",
    category: "Sports",
    price: 300,
    available: true
}
addProduct(Product);
updateProduct(1, { price: 400, available: false, name: "Reference Book" });
filterAvailable();
filterbyCategory("Sports");

