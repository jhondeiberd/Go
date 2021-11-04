class HousePlan{

   constructor(){
       if (this.constructor == HousePlan){
           throw new Error("Abstract classes can't be instantiated.");
       }
   }
  
  setBasement(basement){
    throw new Error("Method 'setBasement()' must be implemented.");
  }

  setStructure(structure){
    throw new Error("Method 'setStructure()' must be implemented.");
  }

  setRoof(roof){
    throw new Error("Method 'setRoof()' must be implemented.");
  }
   
  setInterior(interior){
    throw new Error("Method 'setInterior()' must be implemented.");
  }

  getRoof(){
    throw new Error("Method 'getRoof()' must be implemented.");
  }

  getInterior(){
    throw new Error("Method 'getInterior()' must be implemented.");
  }

  getBasement(){
    throw new Error("Method 'getBasement()' must be implemented.");
  }

  getStructure(){
    throw new Error("Method 'getStructure()' must be implemented.");
  }
}


class House extends HousePlan{

  constructor(){
      super();
  }

  setBasement(basement){
      this.basement = basement;
  }

  setStructure(structure){
      this.structure=structure;
  }

  setInterior(interior){
      this.interior=interior;
  }
  
  setRoof(roof){
      this.roof=roof;
  }

  getInterior(){
      return this.interior;
  }
  
  getRoof(){
      return this.roof;
  }

  getBasement(){
      return this.basement;
  }

  getStructure(){
      return this.structure;
  }
}

//section 2 ----------------------------//

class HouseBuilder{

    //this is the class which has all the steps to create the product
    constructor(){
        if (this.constructor == HouseBuilder){
            throw new Error("this abstract class can not be instantiated.");
        }
    }

    buildBasement(basement){
        throw new Error("method buildBasement() does not have implementation");
    }

    buildStructure(structure){
        throw new Error("method buildStructure() does not have implementation");
    }

    buildInterior(interior){
        throw new Error("method buildInterior() does not have implementation");
    }
    
    buildRoof(roof){
        throw new Error("method buildRoof() does not have implementation");
    }

    GetBaseProduct(){
        throw new Error("method GetBaseProduct() does not have implementation");
    }
}


class BasicHouseBuilder extends HouseBuilder{
    //concrete class one

    constructor(house){
        super();
        this.house = new House();
    }

    buildBasement(){
        return this.house.setBasement("Unfinished basement");
    }

    buildStructure(){
        return this.house.setStructure("Duplex");
    }

    buildInterior(){
        return this.house.setInterior("simple");
    }

    buildRoof(){
        return this.house.setRoof("metal roof");
    }

    GetBaseProduct(){
        return this.house;
    }

}


class HighEndHouseBuilder extends HouseBuilder{
    //concrete class one

    constructor(house){
        super();
        this.house = new House();
    }

    buildBasement(){
        return this.house.setBasement("Indian basement");
    }

    buildStructure(){
        return this.house.setStructure("Single Villa");
    }

    buildInterior(){
        return this.house.setInterior("Maharaja Singh Model");
    }

    buildRoof(){
        return this.house.setRoof("Elon Mask Roofing");
    }

    GetBaseProduct(){
        return this.house;
    }

}


//section 3 -----------------------------------//

class CivilEngineer{

    constructor(houseBuilder){
        this.houseBuilder = houseBuilder;
    }

    constructHouse(){
        //this is the method to control the order of the steps of creation of the product
        this.houseBuilder.buildBasement();
        this.houseBuilder.buildStructure();
        this.houseBuilder.buildRoof();
        this.houseBuilder.buildInterior();
    }

    GetFinalProduct(){
        return this.houseBuilder.GetBaseProduct();
    }

}

//let basicHouseBuilder = new BasicHouseBuilder();
let engineer = new CivilEngineer(new BasicHouseBuilder()); //i just want to remind you that this is AGGREGATION
engineer.constructHouse();

console.log("builder is building structure : " + engineer.GetFinalProduct().getStructure());
console.log("builder is building interior : " + engineer.GetFinalProduct().getInterior());
console.log("builder is building roof : " + engineer.GetFinalProduct().getRoof());
console.log("builder is building basement : " + engineer.GetFinalProduct().getBasement());

console.log('---------------------------------------------------------------------------');

let engineer2 = new CivilEngineer(new HighEndHouseBuilder()); //i just want to remind you that this is AGGREGATION
engineer2.constructHouse();

console.log("builder is building structure : " + engineer2.GetFinalProduct().getStructure());
console.log("builder is building interior : " + engineer2.GetFinalProduct().getInterior());
console.log("builder is building roof : " + engineer2.GetFinalProduct().getRoof());
console.log("builder is building basement : " + engineer2.GetFinalProduct().getBasement());

console.log('---------------------------------------------------------------------------');


