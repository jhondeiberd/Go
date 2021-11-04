/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package HouseBuilder;


public class Main {

	public static void main(String[] args) {
		HouseBuilder basicBuilder = new BasicHouseBuilder();
        CivilEngineer civilEngineer = new CivilEngineer(basicBuilder);
        civilEngineer.constructHouse();
       	System.out.println("Engineer Constructed Basement : " + civilEngineer.getHouse().basement);
       	System.out.println("Engineer Constructed Roof : " + civilEngineer.getHouse().roof);
       	System.out.println("Engineer Constructed Structure : " + civilEngineer.getHouse().structure);
       	System.out.println("Engineer Constructed Interior : " + civilEngineer.getHouse().interior);
       	
       	HouseBuilder highEndBuilder = new HighEndHouseBuilder();
        CivilEngineer civilEngineer2 = new CivilEngineer(highEndBuilder);
        civilEngineer2.constructHouse();
       	System.out.println("Engineer Constructed Basement : " + civilEngineer2.getHouse().basement);
       	System.out.println("Engineer Constructed Roof : " + civilEngineer2.getHouse().roof);
       	System.out.println("Engineer Constructed Structure : " + civilEngineer2.getHouse().structure);
       	System.out.println("Engineer Constructed Interior : " + civilEngineer2.getHouse().interior);
	}

}
