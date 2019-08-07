package com.vng.example.demo;

import org.springframework.http.HttpStatus;
import org.springframework.http.ResponseEntity;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

@RestController
public class HomeController {
    @RequestMapping(value = "/")
    public ResponseEntity<String> home() {
        return new ResponseEntity<>("Hello spring web", HttpStatus.OK);
    }
}
