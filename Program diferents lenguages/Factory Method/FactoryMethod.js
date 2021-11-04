class Shape{
    constructor(){
        if (this.constructor == Shape){
            throw new Error("Abstracct class wont be isntantiated.");
        }
    }

    draw(){
        throw new Error("method draw() must have juice in it");
    }
}

class Square extends Shape{
    constructor(){
        super();
    }

    draw(){
        console.log("inside Square::draw() method.");
    }
}

class Rectangle extends Shape{
    constructor(){
        super();
    }

    draw(){
        console.log("inside Rectangle::draw() method.");
    }
}

class Circle extends Shape{
    constructor(){
        super();
    }

    draw(){
        console.log("inside Circle::draw() method.");
    }
}


//our business logic in order to create the object goes to our factory
//the creation of the object is dont not step by step like bulider pattern,
//but one shot creation. the factory method (ShapeFactory) delegates the rest
//to the sub classes (concrete classes).

class ShapeFactory{

      getShape( shapeType)
    {

        if ( shapeType  == null)
        {
            return null;
        }

        if (shapeType == "CIRCLE")
        {
            return new Circle();
        }
        else if (shapeType == "RECTANGLE")
        {
            return new Rectangle();
        }
        else if (shapeType == "SQUARE")
        {
            return new Square();
        }

        return null;

    };
}



let shapeFactory = new ShapeFactory();

//get a copy of class Circle and invoke the method draw() on it.
let shape1 =   shapeFactory.getShape("CIRCLE");
shape1.draw();


let shape2 = shapeFactory.getShape("RECTANGLE");
shape2.draw();

let shape3 = shapeFactory.getShape("SQUARE");
shape3.draw();
