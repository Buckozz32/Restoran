import React from 'react';
import { useSelector, useDispatch } from 'react-redux';
import { removeFromCart, clearCart } from '../actions';

const Cart = () => {
  const cart = useSelector((state) => state.cart);
  const dispatch = useDispatch();

  const handleRemoveFromCart = (dishId) => {
    dispatch(removeFromCart(dishId));
  };

  const handleClearCart = () => {
    dispatch(clearCart());
  };

  return (
    <div>
      {cart.map((dish) => (
        <div key={dish.id}>
          <h2>{dish.name}</h2>
          <p>{dish.price}</p>
          <button onClick={() => handleRemoveFromCart(dish.id)}>Remove from cart</button>
        </div>
      ))}
      <button onClick={handleClearCart}>Clear cart</button>
    </div>
  );
};

export default Cart;