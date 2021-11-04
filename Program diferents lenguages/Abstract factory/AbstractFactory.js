//section 1 --------------------------//

class TRUCK {
    constructor(){
        if (this.constructor == TRUCK){
            throw new Error("this is an abstract class");
        }
    }
}

class MercedesTRUCK extends TRUCK {
    constructor(){
        super();
    }
}

class VolvoTRUCK extends TRUCK {
    constructor(){
        super();
    }
}

//section 2 ------------------------//
class CAR {
    constructor(){
        if (this.constructor == CAR){
            throw new Error ("this is abstract class");
        }
    }

    display(){
        throw new Error("abstract method display() must be implemented.")
    }
}

class MercedesCAR extends CAR {    
    constructor(){
        super();
    }

    display(){
        console.log("UTTAM car is : MERCEDES 600 AMG");
    }

}

class VolvoCAR extends CAR {    
    constructor(){
        super();
    }

    display(){
        console.log("SAGAR car is : Volvo XC60");
    }

}

//section 3 -----------------------------------//

const MARK = {
    MERCEDES: 1,
    VOLVO: 2
};

Object.freeze(MARK);

//section 4 ------------------------------------//

class AbstractFactory {

      constructor(){
          if (this.constructor == AbstractFactory){
              throw new Error("this is abstract class");
          }
      }

      static getFactory(mark){

          switch(mark){
              case MARK.MERCEDES :
                   return new MercedesDealer();
              case MARK.VOLVO :
                  return new VolvoDealer();
          }

      }

   
    createCAR(){
        throw new Error("method createCar() must be implemented");
    }

    createTRUCK(){
        throw new Error("method createTRUCK() must be implemented");
    }

}



 class VolvoDealer extends AbstractFactory
{

     createCAR()
    {
        return new VolvoCAR();
    }

    createTRUCK()
    {
        return new VolvoTRUCK();
    }

}


 class MercedesDealer extends AbstractFactory
    {

        createCAR()
        {
            return new MercedesCAR();
        }


        createTRUCK()
        {
            return new MercedesTRUCK();
        }

    }



    //section 5 ------------------------------

    let factory = AbstractFactory.getFactory(MARK.VOLVO);
    let volvo = factory.createCAR();
    volvo.display();

    let factory2 = AbstractFactory.getFactory(MARK.MERCEDES);
    let mercedes = factory2.createCAR();
    mercedes.display();