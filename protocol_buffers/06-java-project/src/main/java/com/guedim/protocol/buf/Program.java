package com.guedim.protocol.buf;

import example.simple.Simple.SimpleMessage;


import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.IOException;

public class Program {
    public static void main(String[] args)  {

System.out.println("Hola Mundo");

        SimpleMessage.Builder builder = SimpleMessage.newBuilder();
        builder.setId(42)
                .setIsSimple(true)
                .setName("My simple message name")
                .addSampleList(1)
                .addSampleList(2)
                .addSampleList(3);

        System.out.println(builder.toString());
        SimpleMessage message = builder.build();

        // write protocol to a file
        try {
            System.out.println("Writting to file ...");
            FileOutputStream outputStream = new FileOutputStream("simple-message.bin");
            message.writeTo(outputStream);
            outputStream.close();
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        } catch (IOException e) {
            e.printStackTrace();
        }

        // Read from file
        try {
            System.out.println("Reading from file ...");
            FileInputStream inputStream = new FileInputStream("simple-message.bin");
            SimpleMessage messageFromFile =  SimpleMessage.parseFrom(inputStream);
            System.out.println(messageFromFile.toString());
        } catch (FileNotFoundException e) {
            e.printStackTrace();
        } catch (IOException e) {
            e.printStackTrace();
        }
    }
}
