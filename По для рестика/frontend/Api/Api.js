import axios from 'axios';

const BASE_URL = 'https://my-restaurant-api.com/api';

export const getDishes = async () => {
  try {
    const response = await axios.get(`${BASE_URL}/dishes`);
    return response.data;
  } catch (error) {
    console.error(error);
    throw error;
  }
};

export const addToCart = async (dish) => {
  try {
    const response = await axios.post(`${BASE_URL}/cart`, dish);
    return response.data;
  } catch (error) {
    console.error(error);
    throw error;
  }
};

export const removeFromCart = async (dishId) => {
  try {
    const response = await axios.delete(`${BASE_URL}/cart/${dishId}`);
    return response.data;
  } catch (error) {
    console.error(error);
    throw error;
  }
};

export const clearCart = async () => {
  try {
    const response = await axios.delete(`${BASE_URL}/cart`);
    return response.data;
  } catch (error) {
    console.error(error);
    throw error;
  }
};