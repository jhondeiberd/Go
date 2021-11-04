package HouseBuilder;

public class House implements HousePlan {

	public String basement;
	public String structure;
	public String interior;
	public String roof;
	
	@Override
	public void setBasement(String basement) {
		this.basement = basement;
	}

	@Override
	public void setStructure(String structure) {
		this.structure = structure;
	}

	@Override
	public void setInterior(String interior) {
		this.interior = interior;
	}

	@Override
	public void setRoof(String roof) {
		this.roof = roof;
	}

}
