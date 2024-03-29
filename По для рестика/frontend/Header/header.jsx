import React from 'react';
import { View, Text, TouchableOpacity } from 'react-native';
import { colors, fonts, sizes } from '../style';
import Icon from 'react-native-vector-icons/FontAwesome';

const Header = ({ navigation }) => {
  const handleOpenMenu = () => {
    navigation.openDrawer();
  };

  return (
    <View style={{ flexDirection: 'row', alignItems: 'center', justifyContent: 'space-between', padding: sizes.base, backgroundColor: colors.white }}>
      <TouchableOpacity onPress={handleOpenMenu}>
        <Icon name="bars" size={24} color={colors.gray} />
      </TouchableOpacity>
      <Text style={{ fontFamily: fonts.primary, fontSize: sizes.base * 1.5, color: colors.gray }}>Restaurant Name</Text>
      <View style={{ width: sizes.base * 3 }} />
    </View>
  );
};

export default Header;