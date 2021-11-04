using System;
using System.Collections.Generic;
using System.Text;

namespace HouseBuilder
{
    public interface HousePlan
    {
        void setBasement(String basement);
        void setStructure(String structure);
        void setInterior(String interior);
        void setRoof(String roof);

    }
}
