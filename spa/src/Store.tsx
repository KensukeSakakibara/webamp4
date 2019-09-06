
export interface TaskInterface {
  id: number;
  text: string;
  done: boolean;
}

export interface StoreInterface {
  tasks: TaskInterface[];
}
