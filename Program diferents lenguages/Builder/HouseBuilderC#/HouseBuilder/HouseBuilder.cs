using System;
using System.Collections.Generic;
using System.Text;

namespace HouseBuilder
{
    public interface HouseBuilder
    {
        //this abstract builder defines ALL THE STEPS THAT MUST BE TAKEN
        //in order to CORRECTLY create a SPECIFIC PRODUCT
        void buildRoof();
        void buildInterior();
        void buildBasement();
        void buildStructure();

        House getHouse();

    }
}
