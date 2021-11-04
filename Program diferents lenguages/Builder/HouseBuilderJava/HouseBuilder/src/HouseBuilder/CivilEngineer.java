package HouseBuilder;



public class CivilEngineer
{
    private HouseBuilder houseBuilder;

    public CivilEngineer(HouseBuilder houseBuilder)
    {
        this.houseBuilder = houseBuilder;
    }
    
    public void constructHouse() {
    	this.houseBuilder.buildBasement();
    	this.houseBuilder.buildInterior();
    	this.houseBuilder.buildRoof();
    	this.houseBuilder.buildStructure();
    }
    
    public House getHouse() {
    	return this.houseBuilder.getHouse();
    }
}