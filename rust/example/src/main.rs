use std::f32::consts;
use std::io::{stdin};
mod single_module;
mod sub_module;

static _PI:f32 = 3.14;
static _NAME:&'static str = "Constant";

fn main() {
    println!("Hello world");
    println!("{} of {} abc", 1, 2);
    println!("{number:>0width$}", number=1, width=6);

    let _logical : bool = true;
    let mut mytype = 5;
    println!("First value: {}", mytype);
    mytype = 7;
    println!("Value: {}", mytype);

    println!("PI: {}", consts::PI);

    let _num = 19;
    if _num == 9 {
        static OTHER_NUM:u8 = 10;
        println!("Other num: {}", OTHER_NUM);
    } else {
        println!("No other num");
    }

    let _check = if _num == 19 {true} else { false };
    println!("Check value: {}", _check);

    for i in 1..5 {
        println!("value {}", i);
    }

    let mut index:u8 = 10;
    while index > 0 {
        println!("Value {}", index);
        index = index - 1;
    }

    hello("bill");

    println!("Factorial {}", factorial(5));

    let list_lang:[&str; 2] = ["Rust", "C++"];
    println!("Length: {}", list_lang.len());
    println!("{:?}", list_lang);

    let list_init = ["Init"; 3];
    println!("{:?}", list_init);

    let mut list_num = [10; 10];
    for i in 0..10 {
        list_num[i] = i as u8
    }
    println!("List number {:?}", list_num);

    let mut vec:Vec<u8> = Vec::new();
    let mut another_vec:Vec<u8> = vec![];

    for i in 0..10 {
        vec.push(i as u8);
        another_vec.push(i as u8);
    }

    let sub_array = &another_vec[2..8];

    println!("Vector number {:?}", vec);

    println!("Sub array {:?}", sub_array);

    let sub_array_iter = sub_array.iter();

    for item in sub_array_iter {
        println!("Iterator value {}", item);
    }

    let my_tuple = ("Hello", 1, true);
    println!("My tuple value {:?}", my_tuple);

    let (str, num, check) = my_tuple;

    println!("{}, {}, {}", str, num, check);

    #[allow(dead_code)]
    enum LANGUAGES {Java, C, Rust}
    let _enum_lang = LANGUAGES::Java;

    println!("Input your name: ");
    let mut buf = String::new();
    stdin().read_line(& mut buf).ok().expect("error!");
    println!("Hello {}", buf);

    println!("Input your birth year: ");
    let mut buffer = String::new();
    stdin().read_line(& mut buffer).ok().expect("error!");

    let my_year: Result<u32, _> = buffer.trim().parse();
    match my_year {
        Ok(y) => println!("Birth year {}", y),
        Err(e) => println!("Error {}", e),
    }

    let current_lang : &str = match _enum_lang {
        LANGUAGES::Java => "Java",
        LANGUAGES::C => "C++11",
        LANGUAGES::Rust => "Rust",
    };
    println!("Current language {}", current_lang);

    high_order_func("bill", &hello);

    higher_order_func_with_closure(5, |num| {return 2*num});

    let gen: MyGenericStruct<&str> = MyGenericStruct {
        value: "Generic"
    };
    println!("Generic value: {}", show_generic_func(gen));

    let mut book :Book = Book {
        name: "My book",
        year: 2019,
    };

    book.show();
    book.change_name("Hello book");
    book.show();

    let blog: Blog = Blog{};
    let forum: Forum = Forum{};

    blog.publish();
    forum.publish();
    blog.delete();

    #[allow(dead_code)] // => disable code warning
    struct Structure(i32); // struct like tuple

    struct Person<'a> {
        name: &'a str,
        age: u8,
    }

    let person = Person {
        name: "Bill",
        age: 22,
    };

    println!("Person: name {}, age {}", person.name, person.age);

    #[allow(dead_code)]
    struct Bottle {
        year: i32,
        name: &'static str,
    }

    let my_bottle = Bottle {
        year: 2019,
        name: "Self",
    };

    let mut your_bottle = my_bottle;

//    println!("My bottle {}", my_bottle.name);
    println!("Your bottler {}", your_bottle.name);
    your_bottle.name = "Me";
    println!("Your bottler {}", your_bottle.name);

    let your_bottle_1 = &mut your_bottle;
    your_bottle_1.name = "hi";
    println!("Your bottle {}", your_bottle_1.name);

    your_bottle.name = "hello";

    println!("My bottle {}", your_bottle.name);

    let s1 = String::from("Hello");
    let len = calculate_length(s1.as_str());
    println!("Length of {} is {}", s1, len);

    let len2 = calculate_length_not_ref(s1);
    println!("Len of is {}", len2);

    sub_module::bar::bar();
}

fn calculate_length(str: &str) -> usize {
    str.len()
}

fn calculate_length_not_ref(str: String) -> usize {
    str.len()
}

trait Web {
    fn publish(&self);
    fn delete(&self) {
        println!("Delete web");
    }
}

struct Blog {}

struct Forum {}

impl Web for Blog {
    fn publish(&self) {
        println!("Publish blog");
    }
}

impl Web for Forum {
    fn publish(&self) {
        println!("Publish forum");
    }
}


struct MyGenericStruct<T> {
    value: T,
}

struct Book {
    name: &'static str,
    year: u64,
}

impl Book {
    fn show(&self) {
        println!("Information name {}, year {}", self.name, self.year)
    }
    fn change_name(&mut self, new_name: &'static str) {
        println!("Change name from {} to {}", self.name, new_name);
        self.name = new_name;
    }
}

fn show_generic_func<T>(gen:MyGenericStruct<T>) -> T {
    return gen.value
}

fn hello(name: &str) {
    println!("Hello {}", name)
}

fn factorial(num: u32) -> u32 {
    if num == 1 || num == 0 {
        return 1
    }
    return num * factorial(num - 1)
}

fn high_order_func<MFunc: Fn(&str)>(name:&str, my_func: MFunc) {
    println!("High order function say hi!");
    my_func(name)
}

fn higher_order_func_with_closure<MFunc: Fn(u32) -> (u32)>(num:u32, my_func: MFunc) {
    println!("Higher order function with closure say hi!");
    println!("Result HOF: {}", my_func(num));
}