"use client";

import React, { useState, useEffect, useRef } from "react";
import { BackgroundGradientAnimation } from "../components/ui/background-gradient-animation.tsx";
import { HoverEffect } from "../components/ui/car-hover-effect.tsx";

export default function CardHoverEffectDemo() {
  const [isModalOpen, setIsModalOpen] = useState(false);
  const plus = [
    {
      title: "click to add a new to do item",
      description: "",
      link: "",
    },
  ];
  const modalRef = useRef(null);

  useEffect(() => {
    const handleOutsideClick = (event) => {
      if (modalRef.current && !modalRef.current.contains(event.target)) {
        setIsModalOpen(false);
      }
    };

    if (isModalOpen) {
      document.addEventListener("mousedown", handleOutsideClick);
    }

    return () => {
      document.removeEventListener("mousedown", handleOutsideClick);
    };
  }, [isModalOpen]);

  const toggleModal = () => {
    setIsModalOpen(!isModalOpen);
  };

  return (
    <div className="">
      <BackgroundGradientAnimation>
        <div className="text-5xl lg:text-9xl mt-2 text-center p-3">
          just do it
        </div>
        <div className="text-4xl text-center p-3">
          what are you waiting for?
        </div>
        <div className="max-w-6xl mx-auto px-8 z-10">
          <HoverEffect items={projects} />
        </div>
        <div className="max-w-6xl mx-auto px-8 z-10" onClick={toggleModal}>
          {plus.map((item, index) => (
            <div key={index}>
              <HoverEffect items={plus} />
            </div>
          ))}
        </div>
        <div className="">
          {isModalOpen && (
            <div className="fixed top-0 left-0 w-full h-full flex items-center justify-center z-50">
              <div className="absolute w-full h-full bg-gray-900 opacity-50"></div>
              <div
                ref={modalRef}
                className="bg-zinc-950 rounded-md z-50 overflow-y-auto"
              >
                <div className="max-w-md mx-auto relative overflow-hidden border border-slate-400 border-2 z-10 bg-black p-12 rounded-lg shadow-md before:bg-purple-600 before:rounded-full before:-z-10 before:blur-2xl after:w-32 after:h-32 after:absolute after:bg-sky-400 after:rounded-full after:-z-10 after:blur-3xl after:top-24 after:-right-10">
                  <h2 className="text-2xl font-bold text-white mb-8">
                    Add new to_do item
                  </h2>
                  <form method="post" action="#">
                    <div className="mb-4">
                      <label className="block text-sm font-medium text-gray-300">
                        Title
                      </label>
                      <input
                        className="mt-1 p-2 w-full bg-zinc-900 border border-gray-600 rounded-md text-white"
                        type="text"
                      />
                    </div>
                    <div className="mb-4">
                      <label className="block text-sm font-medium text-gray-300">
                        Description
                      </label>
                      <input
                        className="mt-1 p-2 w-full bg-zinc-900 border border-gray-600 rounded-md text-white"
                        name="description"
                        id="description"
                        type="description"
                      />
                    </div>
                    <div className="mb-4">
                      <label className="block text-sm font-medium text-gray-300">
                        Link
                      </label>
                      <input
                        className="mt-1 p-2 w-full bg-zinc-900 border border-gray-600 rounded-md text-white"
                        rows="3"
                        name="bio"
                        id="bio"
                      ></input>
                    </div>
                    <div className="flex justify-end">
                      <button
                        className="group mt-7 group-hover:before:duration-500 group-hover:after:duration-500 after:duration-500 hover:border-rose-300 hover:before:[box-shadow:_20px_20px_20px_30px_#a21caf] duration-500 before:duration-500 hover:duration-500 underline underline-offset-2 hover:after:-right-8 hover:before:right-12 hover:before:-bottom-8 hover:before:blur hover:underline hover:underline-offset-4  origin-left hover:decoration-2 hover:text-rose-300 relative bg-neutral-800 h-14 w-52 border text-left p-3 text-gray-50 text-base font-bold rounded-lg  overflow-hidden  before:absolute before:w-12 before:h-12 before:content[''] before:right-1 before:top-1 before:z-10 before:bg-violet-500 before:rounded-full before:blur-lg  after:absolute after:z-10 after:w-20 after:h-20 after:content['']  after:bg-rose-300 after:right-8 after:top-3 after:rounded-full after:blur-lg"
                        type="submit"
                      >
                        Submitt
                      </button>
                    </div>
                  </form>
                </div>
              </div>
            </div>
          )}
        </div>
      </BackgroundGradientAnimation>
    </div>
  );
}

export const projects = [
  // fetch from db
  {
    title: "go mini proji",
    description: "micro-services bull",
    link: "",
  },
  {
    title: "os proji",
    description: "ptree impli? seriously?!!",
    link: "",
  },
  {
    title: "cn proji",
    description: "network network, full network",
    link: "https://google.com",
  },
];
