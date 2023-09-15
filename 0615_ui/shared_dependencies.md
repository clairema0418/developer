the app is: ordering machine with virtual customer service staff and menu display

the files we have decided to generate are: index.html, styles.css, scripts.js

Shared dependencies:
1. Exported variables:
   - staffContainer (id of the DOM element containing customer service staff)
   - menuContainer (id of the DOM element containing the menu)
   - cart (array to store selected dishes)
   - totalPrice (variable to store the total price of selected dishes)

2. Data schemas:
   - Dish (object with properties: id, name, description, price, image)

3. Id names of DOM elements:
   - staffContainer (id for the customer service staff container)
   - menuContainer (id for the menu container)
   - dishCard-{id} (id for each dish card, where {id} is the dish id)
   - addToCartButton-{id} (id for the add to cart button on each dish card, where {id} is the dish id)
   - cartSummary (id for the cart summary container)
   - checkoutButton (id for the checkout button)

4. Message names:
   - dishRecommendation (message for recommending and introducing dishes)

5. Function names:
   - displayStaff()
   - displayMenu()
   - addToCart(dishId)
   - removeFromCart(dishId)
   - updateCartSummary()
   - proceedToCheckout()