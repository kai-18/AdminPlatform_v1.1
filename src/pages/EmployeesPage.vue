<template>
  <q-page>
    <q-card>
      <q-card-section>
        <div class="text-h6">User Management</div>
      </q-card-section>

      <q-card-section>
        <q-form @submit.prevent="createEmployee">
          <q-input v-model="newEmployee.name" label="Name" dense />
          <q-input v-model="newEmployee.lastname" label="Lastname" dense />
          <q-input v-model="newEmployee.username" label="Username" dense />
          <q-input v-model="newEmployee.email" label="Email" dense />
          <q-input v-model="newEmployee.position" label="Position" dense />
          <q-input v-model="newEmployee.address" label="Address" dense />
          <q-input v-model="newEmployee.date_of_birth" label="Date of Birth" dense />
          <q-input v-model="newEmployee.place_of_birth" label="Place of Birth" dense />
          <q-btn type="submit" label="Add User" color="primary" class="q-mt-sm" />
        </q-form>
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
              <q-btn
                flat
                dense
                icon="edit"
                @click="openEditDialog(props.row)"
              />
              <q-btn
                flat
                dense
                icon="delete"
                color="negative"
                @click="deleteEmployee(props.row.id)"
              />
            </q-td>
          </template>
        </q-table>
      </q-card-section>
    </q-card>

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
            <q-input v-model="editEmployee.date_of_birth" label="Date of birth (yyyy/mm/dd)" dense />
            <q-input v-model="editEmployee.place_of_birth" label="Place of birth" dense />
            <q-btn type="submit" label="Update" color="primary" class="q-mt-sm" />
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
const newEmployee = ref({
  name: '',
  lastname: '',
  username: '',
  email: '',
  position: '',
  date_of_birth: '',
  place_of_birth: ''
})
const editEmployee = ref({})
const editDialog = ref(false)

const columns = [
  {
    name: 'id',
    required: true,
    label: 'ID',
    align: 'left',
    field: row => row.id,
    format: val => `${val}`,
    sortable: true
  },
  {
    name: 'name',
    required: true,
    label: 'Name',
    align: 'left',
    field: row => row.name, format: val => `${val}`,
    sortable: true
  },
  {
    name: 'lastname',
    align: 'left',
    label: 'Lastname',
    field: row => row.lastname,
    format: val => `${val}`,
    sortable: true
  },
  {
    name: 'username',
    align: 'left',
    label: 'Username',
    field: row => row.username,
    format: val => `${val}`,
    sortable: true
  },
  {
    name: 'email',
    align: 'left',
    label: 'Email',
    field: row => row.email,
    format: val => `${val}`,
    sortable: true
  },
  {
    name: 'position',
    align: 'left',
    label: 'Position',
    field: row => row.position,
    format: val => `${val}`,
    sortable: true
  },
  {
    name:'address',
    align: 'left',
    label: 'Address',
    field: row => row.address,
    format: val => `${val}`
  },
  {
    name: 'dateofbirth',
    align: 'left',
    label: 'Date of Birth',
    field: row => row.date_of_birth,
    format: val => {
      const date = new Date(val)
      return date.toLocaleDateString('en-GB')
    },
    sortable: true
  },
  {
    name: 'placeofbirth',
    align: 'left',
    label: 'Place of Birth',
    field: row => row.place_of_birth,
    format: val => `${val}`,
    sortable: true
  },
  {
    name: 'actions',
    align: 'center',
    label: 'Actions',
    field: 'actions'
  }
]

async function fetchEmployees() {
  try {
    const response = await axios.get('http://localhost:8080/api/employees')
    console.log('Fetched employees:', response.data)
    employees.value = response.data
  } catch (error) {
    console.error('Error fetching employees:', error)
  }
}

async function createEmployee() {
  try {
    await axios.post('http://localhost:8080/api/employees/create', newEmployee.value)
    newEmployee.value = {
      name: '',
      lastname: '',
      username: '',
      email: '',
      position: '',
      address: '',
      date_of_birth: '',
      place_of_birth: ''
    }
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
    await axios.delete('http://localhost:8080/api/employees/delete', {
      data: { id: id }
    })
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
  max-height: 400px;  /* Set height for scrollable container */
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
