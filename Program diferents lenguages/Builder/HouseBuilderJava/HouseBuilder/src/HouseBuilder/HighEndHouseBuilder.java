/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package HouseBuilder;


public class HighEndHouseBuilder implements HouseBuilder{

	private House house;
	
	public HighEndHouseBuilder() {
		this.house = new House();
	}
	
	@Override
	public void buildBasement() {
		house.setBasement("Fully Furnished");
	}

	@Override
	public void buildRoof() {
		house.setRoof("Wooden Roof");
	}

	@Override
	public void buildStructure() {
		house.setStructure("Strongest");
	}

	@Override
	public void buildInterior() {
		house.setInterior("Good");
	}
	
	public House getHouse() {
		return this.house;
		
	}

}

