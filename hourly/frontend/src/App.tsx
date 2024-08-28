import { useQuery } from '@tanstack/react-query'
import { Button, Table,
  Thead,
  Tbody,
  Tr,
  Th,
  Td,
  TableContainer,
  Spinner,
  Heading,
 } from '@chakra-ui/react'
import './App.css'

export type Task = {
  ID: number
  name: string
  project: string
  hours: number
}

const API_URI : string = "http://localhost:8080"

//const API_URL = process.env.API_URL
const handleExportCsv = async(e : React.MouseEvent<HTMLButtonElement>) => {
  e.preventDefault()
    try {
    const res = await fetch( API_URI + "/data/tasks/csv")
    const data = await res.json()
    if (!res.ok) {
      throw new Error(data.error || "Something went wrong")
    } 
    return data
  } catch (error) {
    console.log(error)
  }
}

const App: React.FC = () => {  
  const {data: tasks, isLoading } = useQuery({
    queryKey:["tasks"],
    queryFn: async() => {
      try {
        const res = await fetch(API_URI + "/data/tasks")
        const data = await res.json()
        if (!res.ok) {
          throw new Error(data.error || "Something went wrong")
        }
        return data || [];
      } catch (error) {
        console.log(error)
      }
    }
  })

  return (
    <>
      <nav>
        <ul>
          <li><a href="#">Dashboard</a></li>
          <li><a href="#">Reports</a></li>
          <li><a href="#">About</a></li>
        </ul>
      </nav>
      <div>
        <Heading as='h1' size='2xl' noOfLines={1}>All Tasks</Heading>
        <TableContainer maxWidth={'100%'}>
          <Table variant='striped' colorScheme='grey'>
            <Thead>
              <Tr>
                <Th>ID</Th>
                <Th>Project</Th>
                <Th>Task Name</Th>
                <Th>Hours</Th>
              </Tr>
            </Thead>
            {!isLoading && tasks?.length > 0 && (
              <Tbody>
                {tasks?.map((task : Task) => (
                  <Tr key={task.ID}>
                    <Td>{task.ID}</Td>
                    <Td>{task.project}</Td>
                    <Td>{task.name}</Td>
                    <Td>{task.hours}</Td>
                  </Tr> 
                ))}
              </Tbody>
            )}
          </Table>
        </TableContainer>
        {isLoading && (
          <Spinner />
        )}
        {!isLoading && tasks?.length > 0 && (
          <Button onClick={handleExportCsv}>Export CSV</Button> 
        )}
        {!isLoading && tasks?.length === 0 && (
          <div>No tasks!</div>
        )}
      </div>
      <footer>
        Crafted by Jon 
      </footer>
    </>
  )
}

export default App
