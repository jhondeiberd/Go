using System;
using System.Collections.Generic;
using System.Text;

namespace HouseBuilder
{
    public class House : HousePlan
    {

        //these 4 "small" objects are part of the "big" (complex) object: House
        //we can call this House the complex product

        public String basement;
        public String structure;
        public String roof;
        public String interior;

        //public House() { }  this without parameter constructor is created implicitly for us.

        //the responsibilities of these setters are to configuare the base product to the one we want
        public void setBasement(String basement)
        {
            this.basement = basement;
        }

        public void setStructure(String structure)
        {
            this.structure=structure;
        }

        public void setRoof(String roof)
        {
            this.roof=roof;
        }

        public void setInterior(String interior)
        {
            this.interior = interior;
        }


    }
}





