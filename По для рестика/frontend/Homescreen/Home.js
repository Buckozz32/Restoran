import React, { useEffect, useState } from 'react';
import { View, Text, FlatList, StyleSheet } from 'react-native';
import { getDishes } from '../api';
import DishItem from '../components/DishItem';

const HomeScreen = () => {
  const [dishes, setDishes] = useState([]);

  useEffect(() => {
    const fetchDishes = async () => {
      const data = await getDishes();
      setDishes(data);
    };

    fetchDishes();
  }, []);

  return (
    <View style={styles.container}>
      <FlatList
        data={dishes}
        renderItem={({ item }) => <DishItem dish={item} />}
        keyExtractor={(item) => item.id.toString()}
      />
    </View>
  );
};

const styles = StyleSheet.create({
  container: {
    flex: 1,
    padding: 16,
    backgroundColor: '#fff',
  },
});

export default HomeScreen;