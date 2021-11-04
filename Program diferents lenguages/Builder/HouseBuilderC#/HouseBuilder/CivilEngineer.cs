using System;
using System.Collections.Generic;
using System.Text;

namespace HouseBuilder
{
    class CivilEngineer
    {

        private HouseBuilder houseBuilder;

        //the director decides for the type of the product
        public CivilEngineer(HouseBuilder houseBuilder)
        {
            this.houseBuilder = houseBuilder;
        }


        //the director class controls the algorithm that generating the final product object.
        //in the other words, the correct order of building the complex object.
        public void constructHouse()
        {
            this.houseBuilder.buildBasement();
            this.houseBuilder.buildStructure();
            this.houseBuilder.buildRoof();
            this.houseBuilder.buildInterior();
        }



        public House getHouse()
        {
            return this.houseBuilder.getHouse();
        }


    }
}
