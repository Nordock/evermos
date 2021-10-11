# evermos
evermos

mainm code on the master branch.


Answer:

1. It happened because it was not syncronize between Inventory quantity with the product quantity that customer can be order.
2. THe solution is, when customer add to chart(Create new order) we must check the quantity of the product on the inventory. when we check if quantity inventory is last than zero, we must make a error message "Stock unavailable"
