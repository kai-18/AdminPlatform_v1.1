<template>
  <q-page>
    <q-card>
      <q-card-section>
        <div class="text-h6">Employee Management</div>
      </q-card-section>

      <q-card-section>
        <q-btn label="Add Employee" color="primary" @click="addDialog = true" />
      </q-card-section>

      <q-separator />

      <q-card-section>
        <q-table
          :rows="employees"
          :columns="columns"
          row-key="id"
          flat
          dense
          virtual-scroll
          :rows-per-page-options="[30, 50, 100]"
          class="scrollable-table"
        >
          <template v-slot:body-cell-actions="props">
            <q-td :props="props">
              <q-btn flat dense icon="edit" @click="openEditDialog(props.row)" />
              <q-btn flat dense icon="delete" color="negative" @click="deleteEmployee(props.row.id)" />
            </q-td>
          </template>
        </q-table>
      </q-card-section>
    </q-card>

    <!-- Add Employee Dialog -->
    <q-dialog v-model="addDialog">
      <q-card style="width: 500px;">
        <q-card-section>
          <div class="text-h6">Add New Employee</div>
        </q-card-section>

        <q-card-section>
          <q-form @submit.prevent="createEmployee">
            <q-input v-model="newEmployee.name" label="Name" dense />
            <q-input v-model="newEmployee.lastname" label="Lastname" dense />
            <q-input v-model="newEmployee.username" label="Username" dense />
            <q-input v-model="newEmployee.email" label="Email" dense />
            <q-input v-model="newEmployee.position" label="Position" dense />
            <q-input v-model="newEmployee.address" label="Address" dense />
            <q-input v-model="newEmployee.date_of_birth" label="Date of Birth (yyyy-mm-dd)" dense />
            <q-input v-model="newEmployee.place_of_birth" label="Place of Birth" dense />

            <q-card-actions align="right">
              <q-btn flat label="Cancel" color="negative" @click="addDialog = false" />
              <q-btn type="submit" label="Save" color="primary" />
            </q-card-actions>
          </q-form>
        </q-card-section>
      </q-card>
    </q-dialog>

    <!-- Edit Employee Dialog -->
    <q-dialog v-model="editDialog">
      <q-card>
        <q-card-section>
          <div class="text-h6">Edit Employee</div>
        </q-card-section>

        <q-card-section>
          <q-form @submit.prevent="updateEmployee">
            <q-input v-model="editEmployee.name" label="Name" dense />
            <q-input v-model="editEmployee.lastname" label="Lastname" dense />
            <q-input v-model="editEmployee.username" label="Username" dense />
            <q-input v-model="editEmployee.email" label="Email" dense />
            <q-input v-model="editEmployee.position" label="Position" dense />
            <q-input v-model="editEmployee.address" label="Address" dense />
            <q-input v-model="editEmployee.date_of_birth" label="Date of Birth (yyyy/mm/dd)" dense />
            <q-input v-model="editEmployee.place_of_birth" label="Place of Birth" dense />

            <q-card-actions align="right">
              <q-btn flat label="Cancel" color="negative" @click="editDialog = false" />
              <q-btn type="submit" label="Update" color="primary" />
            </q-card-actions>
          </q-form>
        </q-card-section>
      </q-card>
    </q-dialog>
  </q-page>
</template>


<script setup>
import { ref, onMounted } from 'vue'
import axios from 'axios'

const employees = ref([])
const addDialog = ref(false) // Controls Add Employee Dialog
const editDialog = ref(false) // Controls Edit Employee Dialog

const newEmployee = ref({
  name: '',
  lastname: '',
  username: '',
  email: '',
  position: '',
  address: '',
  date_of_birth: '',
  place_of_birth: ''
})

const editEmployee = ref({})

const columns = [
  {
    name: 'id',
    label: 'ID',
    field: row => row.id,
    sortable: true },
  {
    name: 'name',
    label: 'Name',
    field: row => row.name,
    sortable: true },
  {
    name: 'lastname',
    label: 'Lastname',
    field: row => row.lastname,
    sortable: true },
  {
    name: 'username',
    label: 'Username',
    field: row => row.username,
    sortable: true },
  {
    name: 'email',
    label: 'Email',
    field: row => row.email,
    sortable: true },
  {
    name: 'position',
    label: 'Position',
    field: row => row.position,
    sortable: true },
  {
    name: 'address',
    label: 'Address',
    field: row => row.address },
  {
    name: 'dateofbirth',
    label: 'Date of Birth (yyyy-mm-dd)',
    field: row => row.date_of_birth,
    sortable: true },
  {
    name: 'placeofbirth',
    label: 'Place of Birth',
    field: row => row.place_of_birth,
    sortable: true },
  {
    name: 'actions',
    label: 'Actions',
    field: 'actions' }
]

async function fetchEmployees() {
  try {
    const response = await axios.get('http://localhost:8080/api/employees')
    employees.value = response.data
  } catch (error) {
    console.error('Error fetching employees:', error)
  }
}

async function createEmployee() {
  try {
    await axios.post('http://localhost:8080/api/employees/create', newEmployee.value)
    newEmployee.value = { name: '', lastname: '', username: '', email: '', position: '', address: '', date_of_birth: '', place_of_birth: '' }
    addDialog.value = false
    fetchEmployees()
  } catch (error) {
    console.error('Error creating employee:', error)
  }
}

function openEditDialog(employee) {
  editEmployee.value = { ...employee }
  editDialog.value = true
}

async function updateEmployee() {
  try {
    await axios.put('http://localhost:8080/api/employees/update', editEmployee.value)
    editDialog.value = false
    fetchEmployees()
  } catch (error) {
    console.error('Error updating employee:', error)
  }
}

async function deleteEmployee(id) {
  try {
    await axios.delete('http://localhost:8080/api/employees/delete', { data: { id: id } })
    fetchEmployees()
  } catch (error) {
    console.error('Error deleting employee:', error)
  }
}

onMounted(() => {
  fetchEmployees()
})
</script>

<style scoped>
.scrollable-table {
  max-height: 600px;  /* Set height for scrollable container */
  overflow-y: auto;   /* Allow vertical scrolling */
}

.q-table__header {
  position: sticky;  /* Sticky header */
  top: 0;
  background-color: white;
  z-index: 1;  /* Ensures header is always on top of the body */
}

.q-th {
  background-color: #f5f5f5;  /* Optional: header background color */
}

.q-mt-sm {
  margin-top: 8px;
}
</style>
