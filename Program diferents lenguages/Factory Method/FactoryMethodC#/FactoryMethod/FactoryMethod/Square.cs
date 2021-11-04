using System;
using System.Collections.Generic;
using System.Text;

namespace FactoryMethod
{
    class Square : Shape
    {
        public void draw()
        {
            Console.WriteLine("inside Square::draw() method.");
        }
    }
}
