import React, { useEffect, useState } from 'react';
import api from '../helpers/api';
import { useParams } from 'react-router-dom';
import { type Task } from './Home';

const TaskDetails : React.FC = () => {
    const { taskId } = useParams()
    const [task, setTask] = useState<Task | null>(null)
    
    useEffect(() => {
        try {
            api.get("http://localhost:8080/data/tasks/" + taskId)
                .then(res => {
                setTask(res.data)
            })
        }
        catch (error) {
            console.error(error)
        }
    }, [taskId])
    
    return (
        <div>
            <h1>Task Details</h1>
            {task && (
                <div>
                <h2>{task.name}</h2>
                <p>{task.project}</p>
                <p>{task.hours}</p>
                <p>{task.target_date}</p>
                </div>
            )}
        </div>
    )
}

export default TaskDetails