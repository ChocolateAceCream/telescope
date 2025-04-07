import { Sketch } from './sketch'
export interface Project {
  id: number
  project_name: string
  comment: string
  updated_at: string
  status: string
  address: string
}

export interface ProjectDetails {
  project: Project
  sketches: Sketch[]
}
