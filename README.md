# Go-funding
Introductory golang example to learn go

This example is based on the tutorial in: https://www.toptal.com/go/go-programming-a-step-by-step-introductory-tutorial

## General

A wrapper which can turn any data structure into a transactional service. We’ll use a Fund type as an example – a simple store for our startup’s remaining funding, where we can check the balance and make withdrawals.

In this introduction to programming in Go, we’ll build the service in small steps, making a mess along the way and then cleaning it up again. Along the way, we’ll encounter lots of cool Go features, including:

    - Struct types and methods
    - Unit tests and benchmarks
    - Goroutines and channels
    - Interfaces and dynamic typing

## A Simple Fund

Let’s write some code to track our startup’s funding. The fund starts with a given balance, and money can only be withdrawn (we’ll figure out revenue later).
