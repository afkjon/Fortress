import React, { useState } from 'react'
import { useQuery } from '@tanstack/react-query'
import { 
  Box,
  Flex,
  Button, 
  Table,
  Thead,
  Tbody,
  Tr,
  Th,
  Td,
  TableContainer,
  Input,
  Spinner,
  Heading,
} from '@chakra-ui/react'

import api from '../helpers/api'

import { Link } from 'react-router-dom'

export type Task = {
  id: number
  name: string
  project: string
  hours: number
  target_date: string
}

const API_URI : string = "http://localhost:8080"

const handleExportCsv = async(e : React.MouseEvent<HTMLButtonElement>) => {
  e.preventDefault()
    try {
    const res = await api.get( API_URI + "/data/tasks/csv")
    if (res.status !== 200) {
      throw new Error(res.data.error || "Something went wrong")
    } 
    return res.data
  } catch (error) {
    console.error(error)
  }
}

interface SearchInput {
    searchInput: string
}

const Home: React.FC = () => {  
  const [searchInput, setSearchInput] = useState<SearchInput>({searchInput: ""})
  const [filteredTasks, setFilteredTasks] = useState<Task[]>([])
  
  const {data: tasks, isLoading } = useQuery({
    queryKey:["tasks"],
    queryFn: async() => {
      try {
        const res = await api.get(API_URI + "/data/tasks")
        if (res.status !== 200) {
          throw new Error("Something went wrong")
        }
        if (res.data.length === 0) {
          return []
        }        
        setFilteredTasks(res.data)
        return res.data;
      } catch (error) {
        console.error(error)
      }
    }
  })

  const handleSearchInput = (e : React.ChangeEvent<HTMLInputElement>) => {
    if (e.target.value === "") {
      setFilteredTasks(tasks)
    } else {
      setFilteredTasks(tasks?.filter((task: { name: string }) => task.name.toLowerCase().includes(e.target.value.toLowerCase())))
    }

    setSearchInput({searchInput: e.target.value})
  }

  return (
    <Box>
      <Heading as='h1' size='2xl' noOfLines={1} my='1rem'>All Tasks</Heading>
      {!isLoading && tasks?.length > 0 && (
          <Flex gap='5px'>
              <Button onClick={handleExportCsv}>Export CSV</Button> 
              <Button onClick={handleExportCsv}>Delete List</Button>
              <Input value={ searchInput.searchInput} 
                  onChange={handleSearchInput} 
                  placeholder='Task Name' 
                  size='md' 
                  />
          </Flex>
      )}
      <TableContainer maxWidth={'100%'}>
          <Table variant='striped'>
          <Thead>
              <Tr>
              <Th>Project</Th>
              <Th>Task Name</Th>
              <Th>Hours</Th>
              <Th>Date</Th>
              </Tr>
          </Thead>
          {!isLoading && filteredTasks?.length > 0 && (
              <Tbody>
              {tasks?.map((task : Task) => (
                  <Tr key={task.id}>
                    <Td>{task.project}</Td>
                    <Td>
                        <Link to={`tasks/` + task.id}>{task.name}</Link>
                    </Td>
                    <Td>{task.hours}</Td>
                    <Td>{task.target_date}</Td>
                    </Tr> 
              ))}
              </Tbody>
          )}
          </Table>
          </TableContainer>
      {isLoading && (
          <Spinner />
      )}
      {!isLoading && tasks?.length === 0 && (
          <div>No tasks!</div>
      )}
    </Box>
  )
}

export default Home
