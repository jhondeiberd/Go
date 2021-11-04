using System;
using System.Collections.Generic;
using System.Text;

namespace HouseBuilder
{
    class BasicHouseBuilder : HouseBuilder
    {
        private House house;

        public BasicHouseBuilder()
        {
            this.house = new House();  //composition
        }


        public void buildBasement()
        {
            house.setBasement("unfinished");
        }

        public void buildStructure()
        {
            house.setStructure("Minimalist");
        }

        public void buildRoof()
        {
            house.setRoof("Metal Roofing");
        }

        public void buildInterior()
        {
            house.setInterior("good");
        }

        public House getHouse()
        {
            return this.house;
        }

    }
}
