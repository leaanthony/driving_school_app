<template>
  <div id="mainMenu">
    <nav class="d-none d-md-block border-right">
      <div class="sidebar-sticky">
        <ul class="nav flex-column">
          <div class="drivingSchoolName">
            <h5>Auto École Noureddine</h5>
          </div>
          <!-- TODO make better solution -->
          <!-- Students list -->
          <li class="nav-item">
            <router-link :to="{ name: 'dashboard' }" class="nav-link">
              Burreau
            </router-link>
          </li>
          <li class="nav-item">
            <router-link :to="{ name: 'listStudents' }" class="nav-link">
              Tous Étudiants
            </router-link>
          </li>
          <li class="nav-item">
            <router-link :to="{ name: 'addStudent' }" class="nav-link">
              Ajouté Un Étudiant
            </router-link>
          </li>
          <li class="nav-item">
            <router-link :to="{ name: 'addExamList' }" class="nav-link">
              Ajouté Une Liste Éxamen
            </router-link>
          </li>
          <li class="nav-item">
            <router-link :to="{ name: 'examLists' }" class="nav-link">
              Toutes Les Liste Éxamen
            </router-link>
          </li>
        </ul>
      </div>
    </nav>
  </div>
</template>

<script>

export default {
  name: "mainMenu",
  data: () => ({
    currentPage: 0,
    items:[],
    pages: [
      {"title": "Students","route": ""},
      {"title": "Add Student", "route": "add-student"},
      {"title": "Add Exam List", "route": "add-exam-list"},
      {"title": "Exam Lists", "route": "exam-lists"}
    ],
  }),
  methods:{
    setCurrentPage: function(index){
      this.currentPage = index;
    }
  }
}
</script>

<style scoped>
nav{
  height: 100%;
}
.drivingSchoolName{
  margin: 0px;
  padding: 25px;
  color: black;
}
.nav-link{
  padding-top: 20px;
  padding-bottom: 20px;
}

a{
  margin: 0px;
  color: black;
}
a:hover{
  color: #0486ff;
  background-color: #e3e3e3;
}
.active{
  border-right: 4px solid #0486ff;
  background-color: #e3e3e3;
  color: #0486ff;
}

</style>