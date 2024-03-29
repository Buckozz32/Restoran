import React, { useDispatch, useSelector } from 'react-redux';
import { View, Text, StyleSheet, FlatList } from 'react-native';
import { addToCart } from '../actions';

const DishList = () => {
  const dishes = useSelector((state) => state.dishes);
  const dispatch = useDispatch();

  const handleAddToCart = (dish) => {
    dispatch(addToCart(dish));
  };

  const renderItem = ({ item }) => (
    <View style={styles.dishItem}>
      <Text style={styles.dishName}>{item.name}</Text>
      <Text style={styles.dishPrice}>{item.price} $</Text>
      <Text style={styles.dishDescription}>{item.description}</Text>
      <Button title="Add to cart" onPress={() => handleAddToCart(item)} color="#00BFFF" />
    </View>
  );

  return (
    <View style={styles.container}>
      <FlatList
        data={dishes}
        renderItem={renderItem}
        keyExtractor={(item) => item.id.toString()}
      />
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    padding: 16,
  },
  dishItem: {
    backgroundColor: '#F2F2F2',
    borderRadius: 8,
    padding: 16,
    marginBottom: 16,
  },
  dishName: {
    fontSize: 18,
    fontWeight: 'bold',
    color: '#333',
    marginBottom: 8,
  },
  dishPrice: {
    fontSize: 16,
    color: '#00BFFF',
    marginBottom: 8,
  },
  dishDescription: {
    fontSize: 14,
    color: '#666',
  },
});

export default DishList;