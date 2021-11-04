package HouseBuilder;



public interface HouseBuilder {

	void buildBasement();
	void buildRoof();
	void buildStructure();
	void buildInterior();
	
	House getHouse();
}
