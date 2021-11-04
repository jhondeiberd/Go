
package HouseBuilder;

public class BasicHouseBuilder implements HouseBuilder{

	private House house;
	
	public BasicHouseBuilder() {
		this.house = new House();
	}
	
	@Override
	public void buildBasement() {
		house.setBasement("Unfurnished");
	}

	@Override
	public void buildRoof() {
		house.setRoof("Metal Roof");
	}

	@Override
	public void buildStructure() {
		house.setStructure("Minimal");
	}

	@Override
	public void buildInterior() {
		house.setInterior("Average");
	}
	
	public House getHouse() {
		return this.house;
		
	}

}
