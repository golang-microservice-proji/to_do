import { cn } from "@/utils/cn";
import { AnimatePresence, motion } from "framer-motion";
import { useState } from "react";

export const HoverEffect = ({
  items,
  updateTask,
  deleteTask,
  className,
}: {
  items: {
    ID: string;
    Title: string;
    Deadline: string;
    Completed: boolean;
  }[];
  className?: string;
  updateTask: (id: string) => Promise<void>;
  deleteTask: (id: string) => Promise<void>;
}) => {
  let [hoveredIndex, setHoveredIndex] = useState<number | null>(null);

  // Function to format datetime to human-readable format
  const formatDateTime = (dateTimeString: string) => {
    const date = new Date(dateTimeString);
    const year = date.getUTCFullYear();
    const month = ("0" + (date.getUTCMonth() + 1)).slice(-2);
    const day = ("0" + date.getUTCDate()).slice(-2);
    const hours = ("0" + date.getUTCHours()).slice(-2);
    const minutes = ("0" + date.getUTCMinutes()).slice(-2);
    const seconds = ("0" + date.getUTCSeconds()).slice(-2);
    return `${year}-${month}-${day} ${hours}:${minutes}:${seconds} UTC`;
  };

  //items = items.filter((item) => item.Completed == false);
  return (
    <div
      className={cn(
        "grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 py-10",
        className,
      )}
    >
      {items.map((item, idx) => (
        <div
          key={idx}
          className="relative group block p-2 h-full w-full"
          onMouseEnter={() => setHoveredIndex(idx)}
          onMouseLeave={() => setHoveredIndex(null)}
        >
          <AnimatePresence>
            {hoveredIndex === idx && (
              <motion.span
                className="absolute inset-0 h-full w-full bg-neutral-200 bg-slate-800/[0.8] block rounded-3xl"
                layoutId="hoverBackground"
                initial={{ opacity: 0 }}
                animate={{
                  opacity: 1,
                  transition: { duration: 0.15 },
                }}
                exit={{
                  opacity: 0,
                  transition: { duration: 0.15, delay: 0.2 },
                }}
              />
            )}
          </AnimatePresence>
          <Card>
            <CardTitle>{item.Title}</CardTitle>
            <CardDescription>{formatDateTime(item.Deadline)}</CardDescription>
            {!item.Completed ? (
              <CardMarkCompleted updateFunction={() => updateTask(item.ID)} />
            ) : (
              <CardDelete deleteFunction={() => deleteTask(item.ID)} />
            )}
          </Card>
        </div>
      ))}
    </div>
  );
};
const CardDelete = ({ deleteFunction }: { deleteFunction: () => void }) => {
  return (
    <div className="grid grid-cols-2 pt-5 ">
      <button className=" mr-2 bg-gradient-to-r from-teal-500 to-green-600 text-white font-bold py-3 px-6 rounded-full shadow-lg transform transition-all duration-500 ease-in-out">
        Done!
      </button>
      <button
        className="ml-2 inline-block rounded-full border-2 border-rose-500 text-rose-500 hover:border-rose-600 hover:bg-rose-400 hover:bg-opacity-10 hover:text-rose-600 focus:border-rose-700 focus:text-rose-700 active:border-rose-800 active:text-rose-800 dark:border-rose-300 dark:text-rose-300 dark:hover:hover:bg-rose-300 px-4 pb-2 pt-2.5 text-xs font-medium uppercase leading-normal transition duration-150 ease-in-out focus:outline-none focus:ring-0"
        type="button"
        onClick={deleteFunction}
      >
        Delete
      </button>
    </div>
  );
};
const CardMarkCompleted = ({
  updateFunction,
}: {
  updateFunction: () => void;
}) => {
  return (
    <button
      className="mt-3 p-3 bg-gradient-to-r from-purple-400 to-blue-500 hover:from-pink-500 hover:to-purple-600 text-white font-bold py-3 px-6 rounded-full shadow-lg transform transition-all duration-500 ease-in-out hover:scale-110 hover:brightness-110 hover:animate-pulse active:animate-bounce "
      type="submit"
      onClick={updateFunction}
    >
      Mark as completed
    </button>
  );
};

export const HoverEffectAddToDo = ({
  items,
  className,
}: {
  items: {
    Title: string;
  }[];
  className?: string;
}) => {
  let [hoveredIndex, setHoveredIndex] = useState<number | null>(null);

  return (
    <div
      className={cn(
        "grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 py-10",
        className,
      )}
    >
      {items.map((item, idx) => (
        <div
          key={idx}
          className="relative group block p-2 h-full w-full"
          onMouseEnter={() => setHoveredIndex(idx)}
          onMouseLeave={() => setHoveredIndex(null)}
        >
          <AnimatePresence>
            {hoveredIndex === idx && (
              <motion.span
                className="absolute inset-0 h-full w-full bg-neutral-200 bg-slate-800/[0.8] block rounded-3xl"
                layoutId="hoverBackground"
                initial={{ opacity: 0 }}
                animate={{
                  opacity: 1,
                  transition: { duration: 0.15 },
                }}
                exit={{
                  opacity: 0,
                  transition: { duration: 0.15, delay: 0.2 },
                }}
              />
            )}
          </AnimatePresence>
          <Card>
            <CardTitle>{item.Title}</CardTitle>
          </Card>
        </div>
      ))}
    </div>
  );
};

export const Card = ({
  className,
  children,
}: {
  className?: string;
  children: React.ReactNode;
}) => {
  return (
    <div
      className={cn(
        "rounded-2xl h-full w-full p-4 overflow-hidden bg-black border border-transparent  group-hover:border-slate-700 relative z-20",
        className,
      )}
    >
      <div className="relative z-50">
        <div className="p-4">{children}</div>
      </div>
    </div>
  );
};

export const CardTitle = ({
  className,
  children,
}: {
  className?: string;
  children: React.ReactNode;
}) => {
  return (
    <h4 className={cn("text-zinc-100 font-bold tracking-wide mt-4", className)}>
      {children}
    </h4>
  );
};

export const CardDescription = ({
  className,
  children,
}: {
  className?: string;
  children: React.ReactNode;
}) => {
  return (
    <p
      className={cn(
        "mt-8 text-zinc-400 tracking-wide leading-relaxed text-sm",
        className,
      )}
    >
      {children}
    </p>
  );
};
