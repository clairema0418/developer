const staffContainer = document.getElementById("staffContainer");
const menuContainer = document.getElementById("menuContainer");
const cartSummary = document.getElementById("cartSummary");
const checkoutButton = document.getElementById("checkoutButton");

let cart = [];
let totalPrice = 0;

class Dish {
  constructor(id, name, description, price, image) {
    this.id = id;
    this.name = name;
    this.description = description;
    this.price = price;
    this.image = image;
  }
}

function displayStaff() {
  staffContainer.innerHTML = '<img src="images/virtual_staff.png" alt="Virtual Customer Service Staff">';
}

function displayMenu() {
  const dishes = [
    new Dish(1, "Dish 1", "Description 1", 10, "images/dish1.jpg"),
    new Dish(2, "Dish 2", "Description 2", 15, "images/dish2.jpg"),
    new Dish(3, "Dish 3", "Description 3", 20, "images/dish3.jpg"),
    new Dish(4, "Dish 4", "Description 4", 25, "images/dish4.jpg"),
    new Dish(5, "Dish 5", "Description 5", 30, "images/dish5.jpg"),
    new Dish(6, "Dish 6", "Description 6", 35, "images/dish6.jpg"),
  ];

  dishes.forEach(dish => {
    const dishCard = document.createElement("div");
    dishCard.id = `dishCard-${dish.id}`;
    dishCard.classList.add("dish-card");

    dishCard.innerHTML = `
      <img src="${dish.image}" alt="${dish.name}">
      <h3>${dish.name}</h3>
      <p>${dish.description}</p>
      <p>Price: $${dish.price}</p>
      <button id="addToCartButton-${dish.id}" onclick="addToCart(${dish.id})">Add to Cart</button>
    `;

    menuContainer.appendChild(dishCard);
  });
}

function addToCart(dishId) {
  const dish = document.getElementById(`dishCard-${dishId}`);
  const dishData = {
    id: dishId,
    name: dish.querySelector("h3").textContent,
    price: parseFloat(dish.querySelector("p:last-child").textContent.replace("Price: $", ""))
  };

  cart.push(dishData);
  totalPrice += dishData.price;
  updateCartSummary();
}

function removeFromCart(dishId) {
  const dishIndex = cart.findIndex(dish => dish.id === dishId);
  if (dishIndex !== -1) {
    totalPrice -= cart[dishIndex].price;
    cart.splice(dishIndex, 1);
    updateCartSummary();
  }
}

function updateCartSummary() {
  cartSummary.innerHTML = `Total: $${totalPrice.toFixed(2)}`;
}

function proceedToCheckout() {
  if (cart.length === 0) {
    alert("Please add items to your cart before proceeding to checkout.");
    return;
  }

  // Implement checkout logic here
  console.log("Proceeding to checkout with the following items:", cart);
}

displayStaff();
displayMenu();