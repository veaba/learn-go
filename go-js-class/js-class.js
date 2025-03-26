class Person {
  constructor(name, age) {
    this.name = name;
    this.age = age;
  }

  greet() {
    console.log(`Hello, my name is ${this.name} and I'm ${this.age} years old.`);
  }

  haveBirthday() {
    this.age++;
    console.log(`Happy birthday! Now I'm ${this.age}.`);
  }

  static info() {
    console.log("This is a Person class");
  }
}

// 使用示例
const john = new Person("John", 30);
john.greet();
john.haveBirthday();
Person.info();