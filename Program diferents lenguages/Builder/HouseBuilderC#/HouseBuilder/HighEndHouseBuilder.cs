using System;
using System.Collections.Generic;
using System.Text;

namespace HouseBuilder
{
    class HighEndHouseBuilder : HouseBuilder
    {
        private House house;

        public HighEndHouseBuilder()
        {
            this.house = new House();  //composition
        }


        public void buildBasement()
        {
            house.setBasement("Indian Basement");
        }

        public void buildStructure()
        {
            house.setStructure("Super Single Deluxe");
        }

        public void buildRoof()
        {
            house.setRoof("Solar Tiles Roofing");
        }

        public void buildInterior()
        {
            house.setInterior("Contemporary Maharaja");
        }

        public House getHouse()
        {
            return this.house;
        }

    }
}
